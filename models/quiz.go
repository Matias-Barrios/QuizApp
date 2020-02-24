package models

// Question :
type Question struct {
	ID       string   `json:"id"`
	Mode     string   `json:"mode"`
	Question string   `json:"question"`
	Image    string   `json:"image"`
	Options  []string `json:"options"`
	Answers  []string `json:"answers"`
}

// Quiz :
type Quiz struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	Questions   []Question `json:"questions"`
	Completed   bool       `json:"-"`
}
