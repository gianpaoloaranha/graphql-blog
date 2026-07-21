package post

import "github.com/gianpaoloaranha/go-social-network/internal/app/domain"

type Repository interface {
	CreatePost(post *domain.Post) error
	GetPosts() (*[]domain.Post, error)
	GetPostByID(id string) (*domain.Post, error)
	UpdatePost(post *domain.Post) error
	DeletePost(id string) error
}