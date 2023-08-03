package user_usecase

import (
	"context"
	"fmt"

	"github.com/starcharles/sqlc-example/application"
	"github.com/starcharles/sqlc-example/domain/entities"
	domain "github.com/starcharles/sqlc-example/domain/repositories"
	"github.com/starcharles/sqlc-example/domain/value_object"
)

type UserFindInput struct {
	ID value_object.ID
}

type UserFindOutput struct {
	User *entities.User
}

type userFindUseCase struct {
	userRepository domain.IUserRepository
	input          UserFindInput
}

func (u *userFindUseCase) Handle(ctx context.Context) (*UserFindOutput, error) {
	user, err := u.userRepository.FindOne(ctx, u.input.ID)
	fmt.Println(user)

	if err != nil {
		return nil, err
	}

	return &UserFindOutput{
		User: user,
	}, nil
}

func NewUserFindUseCase(input UserFindInput, userRepository domain.IUserRepository) application.IUseCase[*UserFindOutput] {
	return &userFindUseCase{
		userRepository: userRepository,
		input:          input,
	}
}
