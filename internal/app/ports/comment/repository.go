package comment

import "github.com/gianpaoloaranha/go-social-network/internal/app/domain"

type Repository interface {
	CreateComment(comment *domain.Comment) error
	GetCommentByID(id string) (*domain.Comment, error)
	GetCommentsByPostID(postID string) (*[]domain.Comment, error)
	UpdateComment(comment *domain.Comment) error
	DeleteComment(id string) error
}
