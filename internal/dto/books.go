package dto

type CreateBookRequest struct {
	Title       string `json:"title" example:"The Great Gatsby"`
	Author      string `json:"author" example:"F. Scott Fitzgerald"`
	Description string `json:"description" example:"A classic novel set in the 1920s"`
}

type CreateBookResponse struct {
	ID          uint   `json:"id" example:"1"`
	Title       string `json:"title" example:"The Great Gatsby"`
	Author      string `json:"author" example:"F. Scott Fitzgerald"`
	Description string `json:"description" example:"A classic novel set in the 1920s"`
	UserID      uint   `json:"user_id" example:"1"`
}

type BookResponse struct {
	ID          uint   `json:"id" example:"1"`
	Title       string `json:"title" example:"bigane"`
	Author      string `json:"author" example:"kamo"`
	Description string `json:"description" example:"A nice book"`
}

type UpdateBookRequest struct {
	Title       string `json:"title,omitempty" example:"new title"`
	Author      string `json:"author,omitempty" example:"new author"`
	Description string `json:"description,omitempty" example:"updated description"`
}

type UpdateBookResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type DeleteBookResponse struct {
	Message string `json:"message" example:"Book deleted successfully"`
}

type BookErrorResponse struct {
	Error string `json:"error" example:"Unauthorized"`
}
type DeleteErrorResponse struct {
	Error string `json:"error" example:"Book not found or deletion failed"`
}
