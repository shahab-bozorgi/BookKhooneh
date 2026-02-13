package dto

type CreateReviewRequest struct {
	BookID  uint   `json:"book-id" example:1`
	UserID  uint   `json:"user-id" example:1`
	Rating  int    `json:"rating" example:1`
	Comment string `json:"comment" example:"This is a comment"`
}

type CreateReviewResponse struct {
	ID      uint         `json:"id"`
	BookID  uint         `json:"book-id"`
	User    UserResponse `json:"user"`
	Rating  int          `json:"rating"`
	Comment string       `json:"comment"`
}

type Review struct {
	BookID  uint   `json:"book-id" example:"1"`
	UserID  uint   `json:"user-id" example:"1"`
	Rating  int    `json:"rating" example:"1"`
	Comment string `json:"comment" example:"This is a comment"`
}

type GetAllReviewsResponse struct {
	Reviews []Review `json:"reviews"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
