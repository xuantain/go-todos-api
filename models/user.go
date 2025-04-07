package models

type User struct {
	ID        int    `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      string `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
