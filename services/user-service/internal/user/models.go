package user

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}
