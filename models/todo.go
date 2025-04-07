package models

type Todo struct {
	ID          int    `json:"id,string"`
	UserName    string `json:"username"`
	Description string `json:"description"`
	TargetDate  string `json:"targetDate"`
	Done        bool   `json:"done"`
}
