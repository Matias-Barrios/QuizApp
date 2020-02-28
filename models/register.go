package models

// RegisterBody :
type RegisterBody struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CaptchaID int64  `json:"captchaid"`
	Solution  string `json:"solution"`
}

// ChangePasswordBody :
type ChangePasswordBody struct {
	CurrentPassword   string `json:"currentpassword"`
	NewPassword       string `json:"password"`
	RepeatNewPassword string `json:"repeatpassword"`
}

// SendNewPassword :
type SendNewPassword struct {
	Email string `json:"email"`
}

// SendNewPassword :
type RegisterCaptcha struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
}
