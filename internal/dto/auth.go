package dto

type RegisterRequest struct {
	Username string `json:"username" example:"amu"`
	Email    string `json:"email" example:"amu@example.com"`
	Password string `json:"password" example:"12345678"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterErrorResponse struct {
	Error string `json:"error"`
}

type LoginRequest struct {
	Username string `json:"username" example:"amu"`
	Password string `json:"password" example:"12345678"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginErrorResponse struct {
	Error string `json:"error"`
}
