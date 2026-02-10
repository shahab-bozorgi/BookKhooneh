package dto

type LoginRequest struct {
	Username string `json:"username" example:"amu"`
	Password string `json:"password" example:"12345678"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
