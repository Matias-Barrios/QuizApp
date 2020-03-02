package database

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Matias-Barrios/QuizApp/config"
	"github.com/Matias-Barrios/QuizApp/models"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Captchas :
var Captchas map[string]string

func init() {
	envF := config.EnvironmentFetcher{}
	Captchas = make(map[string]string)
	captchasPath, err := envF.GetValue("CAPTCHAS")
	if err != nil {
		log.Println(err.Error())
		log.Fatalln(err.Error())
	}
	fd, err := os.Open(captchasPath)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln(err.Error())
	}
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Println("Captchas file has a wrong format")
			log.Fatalln("Captchas file has a wrong format")
		}
		Captchas[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

}

func getRandomCaptcha() (string, string) {
	for k, v := range Captchas {
		return k, v
	}
	return "", ""
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	var offset int
	for _, v := range strings.Split(label, "##") {
		point := fixed.Point26_6{fixed.Int26_6(x * 30), fixed.Int26_6(y*30 + offset)}
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: basicfont.Face7x13,
			Dot:  point,
		}
		d.DrawString(v)
		offset += 700
	}
}

func createImageForCaptcha() (string, string, error) {
	img := image.NewRGBA(image.Rect(0, 0, 250, 100))
	captchaQ, captchaA := getRandomCaptcha()
	addLabel(img, 20, 30, captchaQ)
	captchapath := "static/captchas/" + config.RandomString(20) + ".png"
	f, err := os.Create(captchapath)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return "", "", err
	}
	return captchapath, captchaA, nil
}

// GenerateCapctha :
func GenerateCapctha(remoteIP string) (int64, string, error) {
	err := DeleteOlderCaptchas()
	if err != nil {
		log.Println(err.Error())
		return 0, "", err
	}
	path, answer, err := createImageForCaptcha()
	res, err := sqlConnection.Exec(`
		INSERT INTO CAPTCHAS 
		(remote_ip, sent_on, captcha_path,challenge_result)
		VALUES(?, ?, ?, ?)	
	`, remoteIP, time.Now().UTC().Unix(), path, answer)
	if err != nil {
		log.Println(err.Error())
		return 0, "", err
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, "", err
	}
	return lastid, path, nil
}

// GetCaptcha :
func GetCaptcha(id int64) (models.Captcha, error) {
	var captcha models.Captcha
	err := sqlConnection.QueryRow(`
		SELECT id,remote_ip, sent_on, captcha_path,challenge_result
		FROM CAPTCHAS
		WHERE id = ?
		AND sent_on > ?
	`, id, time.Now().Add(-10*time.Minute).UTC().Unix()).Scan(&captcha.ID, &captcha.RemoteIP, &captcha.SentOn, &captcha.Path, &captcha.Result)
	if err != nil {
		log.Println(err.Error())
		return models.Captcha{}, err
	}
	return captcha, nil
}

// DeleteCaptcha :
func DeleteCaptcha(id int64) error {
	captcha, err := GetCaptcha(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = os.Remove(captcha.Path)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = sqlConnection.Exec(`
		DELETE FROM CAPTCHAS 
		WHERE id = ?
	`, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// DeleteOlderCaptchas :
func DeleteOlderCaptchas() error {
	_, err := sqlConnection.Exec(`
		DELETE FROM CAPTCHAS
		WHERE sent_on > ?
	`, time.Now().Add(-10*time.Minute).UTC().Unix())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	fileInfo, err := ioutil.ReadDir("static/captchas")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	now := time.Now()
	for _, info := range fileInfo {
		if diff := now.Sub(info.ModTime()); diff > (10 * time.Minute) {
			err := os.Remove("static/captchas/" + info.Name())
			if err != nil {
				log.Println(err.Error())
				return err
			}
		}
	}
	return nil
}
