package structs

import (
	"time"

	"github.com/google/uuid"
)

type CreatePost struct {
	Title string
	Body  string
}

type UpdatePost struct {
	Title string
	Body  string
}

type Post struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	AuthorID   uuid.UUID `json:"authorId"`
	AuthorName string    `json:"authorName"`
}
