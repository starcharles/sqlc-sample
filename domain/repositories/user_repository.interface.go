package repositories

import (
	"github.com/starcharles/sqlc-example/domain/entities"
	"github.com/starcharles/sqlc-example/domain/value_object"
)

type IUserRepository interface {
	FindOne(id value_object.ID) (*entities.User, error)
}
