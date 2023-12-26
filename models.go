package models

// User struct represents user information
type User struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}
