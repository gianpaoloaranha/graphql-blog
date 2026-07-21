package usecase

import (
	"github.com/gianpaoloaranha/go-social-network/internal/app/domain"
	"github.com/gianpaoloaranha/go-social-network/internal/app/ports/user"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUserUsecase(userRepository user.Repository) user.UseCase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) CreateUser(user user.CreateUserInput) error {
	panic("implement me")
}

func (uc *userUsecase) GetUsers() (*[]domain.User, error) {
	panic("implement me")
}

func (uc *userUsecase) GetUserByID(id string) (*domain.User, error) {
	panic("implement me")
}

func (uc *userUsecase) UpdateUser(user user.UpdateUserInput) error {
	panic("implement me")
}

func (uc *userUsecase) DeleteUser(id string) error {
	panic("implement me")
}

func (uc *userUsecase) FollowUser(followerID, followeeID string) error {
	panic("implement me")
}

func (uc *userUsecase) UnfollowUser(followerID, followeeID string) error {
	panic("implement me")
}

func (uc *userUsecase) GetFollowers(userID string) (*[]domain.User, error) {
	panic("implement me")
}

func (uc *userUsecase) GetFollowing(userID string) (*[]domain.User, error) {
	panic("implement me")
}
