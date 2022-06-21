package models

type User struct {
	// TODO: Implement User model
	ID int64 `json:"id"`
}

type UserRequest struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
