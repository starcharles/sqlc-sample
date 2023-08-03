//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/starcharles/sqlc-example/application"
	user_usecase "github.com/starcharles/sqlc-example/application/user"
	"github.com/starcharles/sqlc-example/infrastracture/database"
)

var providerSet = wire.NewSet(
	// driver
	database.NewConnection,

	// infrastracture
	database.NewUserRepository,
)

func UserFindUseCase(input user_usecase.UserFindInput) (application.IUseCase[*user_usecase.UserFindOutput], error) {
	wire.Build(
		providerSet,
		user_usecase.NewUserFindUseCase,
	)
	return nil, nil
}
