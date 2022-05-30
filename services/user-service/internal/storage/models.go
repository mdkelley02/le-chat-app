package storage

type User struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
