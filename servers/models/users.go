package models

type LoginRequest struct {
	UserName string
	Password string
}

// LoginResponse ...
type LoginResponse struct {
	Token string
}
