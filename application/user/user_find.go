package application

import (
	"time"

	"github.com/starcharles/sqlc-example/domain/repositories"
	"github.com/starcharles/sqlc-example/domain/value_object"
)

type UserFindInput struct {
	ID value_object.ID
}

type UserFindOutput struct {
	ID        value_object.ID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFindUseCase struct {
	userRepository repositories.IUserRepository
}

func (u *UserFindUseCase) Handle(input UserFindInput) (*UserFindOutput, error) {
	user, err := u.userRepository.FindOne(input.ID)

	if err != nil {
		return nil, err
	}

	return &UserFindOutput{
		ID:        user.ID(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}
