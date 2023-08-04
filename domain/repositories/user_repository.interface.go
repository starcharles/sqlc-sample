package domain

import (
	"context"

	"github.com/starcharles/sqlc-example/domain/entities"
	"github.com/starcharles/sqlc-example/domain/value_object"
)

type IUserRepository interface {
	FindOne(context.Context, value_object.ID) (*entities.User, error)
}
