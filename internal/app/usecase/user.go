package usecase

import (
	"net/mail"

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

// CreateUser creates a new user in the system.
func (uc *userUsecase) CreateUser(input user.CreateUserInput) (*domain.User, error) {
	if input.Name == "" {
		return nil, domain.NewError(domain.ErrInvalidInput, "Name is required")
	}

	if input.Email == "" {
		return nil, domain.NewError(domain.ErrInvalidInput, "Email is required")
	}

	if _, err := mail.ParseAddress(input.Email); err != nil {
		return nil, domain.NewError(domain.ErrInvalidInput, "Invalid email format")
	}

	createdUser, err := uc.userRepository.CreateUser(&domain.User{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not create user", err)
	}

	return createdUser, nil
}

// GetUsers retrieves all users from the system.
func (uc *userUsecase) GetUsers() ([]domain.User, error) {
	users, err := uc.userRepository.GetUsers()
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve users", err)
	}

	return users, nil
}

// GetUserByID retrieves a user by their ID.
func (uc *userUsecase) GetUserByID(id string) (*domain.User, error) {
	user, err := uc.userRepository.GetUserByID(id)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	if user == nil {
		return nil, domain.NewError(domain.ErrNotFound, "User not found")
	}

	return user, nil
}

// UpdateUser updates an existing user's information.
func (uc *userUsecase) UpdateUser(user user.UpdateUserInput) (*domain.User, error) {
	userToUpdate, err := uc.userRepository.GetUserByID(user.ID)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	if userToUpdate == nil {
		return nil, domain.NewError(domain.ErrNotFound, "User not found")
	}

	if user.Name != nil {
		userToUpdate.Name = *user.Name
	}

	if user.Email != nil {
		if _, err := mail.ParseAddress(*user.Email); err != nil {
			return nil, domain.NewError(domain.ErrInvalidInput, "Invalid email format")
		}

		userToUpdate.Email = *user.Email
	}

	updatedUser, err := uc.userRepository.UpdateUser(userToUpdate)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not update user", err)
	}
	
	return updatedUser, nil
}

// DeleteUser deletes a user from the system by their ID.
func (uc *userUsecase) DeleteUser(id string) error {
	_, err := uc.userRepository.GetUserByID(id)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	err = uc.userRepository.DeleteUser(id)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not delete user", err)
	}

	return nil
}

// FollowUser allows a user to follow another user.
func (uc *userUsecase) FollowUser(followerID, followeeID string) error {
	_, err := uc.userRepository.GetUserByID(followerID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	_, err = uc.userRepository.GetUserByID(followeeID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	err = uc.userRepository.FollowUser(followerID, followeeID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not follow user", err)
	}

	return nil
}

// UnfollowUser allows a user to unfollow another user.
func (uc *userUsecase) UnfollowUser(followerID, followeeID string) error {
	_, err := uc.userRepository.GetUserByID(followerID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	_, err = uc.userRepository.GetUserByID(followeeID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	err = uc.userRepository.UnfollowUser(followerID, followeeID)
	if err != nil {
		return domain.WrapError(domain.ErrInternal, "Could not follow user", err)
	}

	return nil
}

// GetFollowers retrieves the list of followers for a given user.
func (uc *userUsecase) GetFollowers(userID string) ([]domain.User, error) {
	_, err := uc.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	followers, err := uc.userRepository.GetFollowers(userID)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve followers", err)
	}

	return followers, nil
}

// GetFollowing retrieves the list of users that a given user is following.
func (uc *userUsecase) GetFollowing(userID string) ([]domain.User, error) {
	_, err := uc.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve user", err)
	}

	following, err := uc.userRepository.GetFollowing(userID)
	if err != nil {
		return nil, domain.WrapError(domain.ErrInternal, "Could not retrieve followers", err)
	}

	return following, nil
}
