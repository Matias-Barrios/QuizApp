package models

// RegisterBody :
type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ChangePasswordBody :
type ChangePasswordBody struct {
	CurrentPassword   string `json:"currentpassword"`
	NewPassword       string `json:"password"`
	RepeatNewPassword string `json:"repeatpassword"`
}
