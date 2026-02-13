package dto

type GetUserRequest struct {
	Username string `json:"username" example:"amu"`
}

type UserResponse struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"amu"`
	Email    string `json:"email" example:"amu@example.com"`
}

type GetUserResponse struct {
	User UserResponse `json:"user"`
}

type AllUsersResponse struct {
	Username string `json:"username" example:"amu"`
	Email    string `json:"email" example:"amu@example.com"`
}

type GetAllUsersResponse struct {
	Users []AllUsersResponse `json:"users"`
}

type UserErrorResponse struct {
	Error string `json:"error"`
}
