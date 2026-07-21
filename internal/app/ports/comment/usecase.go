package comment

import "github.com/gianpaoloaranha/go-social-network/internal/app/domain"

type CreateCommentInput struct {
	AuthorID string
	PostID   string
	Message  string
}

type UpdateCommentInput struct {
	ID      string
	Message *string
}

type UseCase interface {
	CreateComment(comment CreateCommentInput) error
	GetCommentByID(id string) (*domain.Comment, error)
	GetCommentsByPostID(postID string) (*[]domain.Comment, error)
	UpdateComment(comment UpdateCommentInput) error
	DeleteComment(id string) error
}
