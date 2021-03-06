package router

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/Matias-Barrios/QuizApp/config"
	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
	"github.com/Matias-Barrios/QuizApp/views"
	"github.com/dgrijalva/jwt-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Claims :
var Claims models.Claim

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := database.GetUser(password, username)
	if err != nil {
		log.Println(err.Error())
		err = database.Log(r.RemoteAddr, username, time.Now().UTC().Unix(), "UNSUCCESSFULLOGINATTEMPT", err.Error())
		if err != nil {
			log.Println(err.Error())
		}
		http.Redirect(w, r, "/login?error=1", 302)
		return
	}
	user.Password = config.RandomString(10)
	claims := &models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(800 * time.Minute).Unix(),
		},
	}
	err = database.Log(r.RemoteAddr, user.Email, time.Now().UTC().Unix(), "LOGIN", "User successfully logged in.")
	if err != nil {
		log.Println(err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 800),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index", 302)
	return
}

func getClaims(w http.ResponseWriter, r *http.Request) models.Claim {
	token, err := r.Cookie("token")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return models.Claim{}
	}
	claims := &models.Claim{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		//log.Println(err.Error())
		return APP_KEY, nil
	})
	return *claims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cerr := database.Log(r.RemoteAddr, "", time.Now().UTC().Unix(), "CONNECTIONATTEMPT", "An user has attempted to connect")
		if cerr != nil {
			log.Println(cerr.Error())
		}
		auth, err := r.Cookie("token")
		if err != nil {

			log.Println(err.Error())
			http.Redirect(w, r, "/login", 302)
			return
		}
		token, err := jwt.Parse(auth.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error decrypting token")
				return nil, fmt.Errorf("Internal server error")
			}
			return []byte(APP_KEY), nil
		})
		if err != nil {
			suberr := database.Log(r.RemoteAddr, "", time.Now().UTC().Unix(), "TOKENERROR", err.Error())
			if suberr != nil {
				log.Println(suberr.Error())
			}
			log.Println(err.Error())
			http.Redirect(w, r, "/login", 302)
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}

func forgotHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forgot" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewForgot.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func sendNewPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sendtp" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	sendNewPassword := models.SendNewPassword{}
	err := decoder.Decode(&sendNewPassword)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newPassword := config.RandomString(12)

	err = database.SetNewPassword(sendNewPassword.Email, newPassword)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Changing password for ", sendNewPassword.Email)
	suberr := database.Log(r.RemoteAddr, sendNewPassword.Email, time.Now().UTC().Unix(), "PASSWORDRESET", "User has forgotten the password")
	if suberr != nil {
		log.Println(suberr.Error())
	}
	body := fmt.Sprintf(`
	You have reseted your LinuxQuizapp.com.uy password.
	This is your temporary password: %s
	Remember to change it  as soon as you login!
		    
	Regards!, 
	Linuxquizapp.com.uy
	`, newPassword)
	if err := smtp.SendMail("mail.linuxquizapp.com.uy:25",
		nil,
		"mail.linuxquizapp.com.uy",
		[]string{sendNewPassword.Email},
		[]byte(body)); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = database.Log(r.RemoteAddr, sendNewPassword.Email, time.Now().UTC().Unix(), "PASSWORDRESET", "User has reseted its password.")
	if err != nil {
		log.Println(err.Error())
	}
}

func changepasswordPOSTHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/changepassword" && r.Method != "POST" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	changePasswordBody := models.ChangePasswordBody{}
	err := decoder.Decode(&changePasswordBody)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := getClaims(w, r)
	log.Println("Changing password for ", claims.User.Email)
	suberr := database.Log(r.RemoteAddr, claims.User.Email, time.Now().UTC().Unix(), "PASSWORDCHANGE", "")
	if suberr != nil {
		log.Println(suberr.Error())
	}

	_, err = database.GetUser(changePasswordBody.CurrentPassword, claims.User.Email)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if changePasswordBody.NewPassword != changePasswordBody.RepeatNewPassword {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eightormore, lower, upper, symbol := verifyPassword(changePasswordBody.NewPassword)
	if !eightormore || !lower || !upper || !symbol {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.UpdateUserPassword(claims.User.ID, changePasswordBody.NewPassword)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{ "status" : "success" }`))
	return
}
