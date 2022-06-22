package models

type Login struct {
	Token string `json:"token"`
	ID    int64  `json:"id"`
}

type LoginRequest struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
