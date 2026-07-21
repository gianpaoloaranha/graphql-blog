package domain

import "time"

type Comment struct {
	ID        string
	AuthorID  string
	PostID    string
	Message   string
	CreatedAt time.Time
}