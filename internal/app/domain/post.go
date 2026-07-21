package domain

import "time"

type Post struct {
	ID          string
	AuthorID    string
	Description string
	CreatedAt   time.Time
}