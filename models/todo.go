package models

type Todo struct {
	ID          int    `json:"todo_id,omitempty"`
	UserId      int    `json:"user_id,omitempty"`
	Description string `json:"description"`
	TargetDate  int64  `json:"target_date"`
	Done        bool   `json:"done"`
}
