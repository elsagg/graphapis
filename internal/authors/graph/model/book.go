package model

type Book struct {
	ID       string `json:"id"`
	AuthorID string `json:"authorId"`
}

func (Book) IsEntity() {}
