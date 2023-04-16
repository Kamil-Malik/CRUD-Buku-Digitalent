package dto

// BookDto represent json model for a book
type BookDTO struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}
