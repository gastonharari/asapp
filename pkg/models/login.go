package models

type Login struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
