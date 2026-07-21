package usecase

import (
	"github.com/gianpaoloaranha/go-social-network/internal/app/domain"
	"github.com/gianpaoloaranha/go-social-network/internal/app/ports/comment"
)

type commentUsecase struct {
	commentRepository comment.Repository
}

func NewCommentUsecase(commentRepository comment.Repository) comment.UseCase {
	return &commentUsecase{
		commentRepository: commentRepository,
	}
}

func (uc *commentUsecase) CreateComment(comment comment.CreateCommentInput) error {
	panic("implement me")
}

func (uc *commentUsecase) GetCommentsByPostID(postID string) (*[]domain.Comment, error) {
	panic("implement me")
}

func (uc *commentUsecase) GetCommentByID(id string) (*domain.Comment, error) {
	panic("implement me")
}

func (uc *commentUsecase) UpdateComment(comment comment.UpdateCommentInput) error {
	panic("implement me")
}

func (uc *commentUsecase) DeleteComment(id string) error {
	panic("implement me")
}
