package models

// Captcha :
type Captcha struct {
	ID       int64
	RemoteIP string
	SentOn   int64
	Path     string
	Result   string
}
