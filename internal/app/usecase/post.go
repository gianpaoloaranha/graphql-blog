package usecase

import (
	"github.com/gianpaoloaranha/go-social-network/internal/app/domain"
	"github.com/gianpaoloaranha/go-social-network/internal/app/ports/post"
)

type postUsecase struct {
	postRepository post.Repository
}

func NewPostUsecase(postRepository post.Repository) post.UseCase {
	return &postUsecase{
		postRepository: postRepository,
	}
}

func (uc *postUsecase) CreatePost(post post.CreatePostInput) error {
	panic("implement me")
}

func (uc *postUsecase) GetPosts() (*[]domain.Post, error) {
	panic("implement me")
}

func (uc *postUsecase) GetPostByID(id string) (*domain.Post, error) {
	panic("implement me")
}

func (uc *postUsecase) UpdatePost(post post.UpdatePostInput) error {
	panic("implement me")
}

func (uc *postUsecase) DeletePost(id string) error {
	panic("implement me")
}
