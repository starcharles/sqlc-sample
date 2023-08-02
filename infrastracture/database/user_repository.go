package database

import (
	"github.com/starcharles/sqlc-example/database"
	"github.com/starcharles/sqlc-example/domain/entities"
	"github.com/starcharles/sqlc-example/domain/value_object"
)

type UserRepository struct {
	client *database.Queries
}

func (r UserRepository) FindOne(id value_object.ID) (*entities.User, error) {
	user, err := r.client.GetUser(id.Value())
	if err != nil {
		return nil, err
	}

	return moduleToEntity(user)
}

func moduleToEntity(model *database.User) (*entities.User, error) *entities.User {
	id, err := value_object.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}

	return entities.NewUser(
		id,
		model.Name,
		model.Email,
		model.Password,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func NewUserRepository(db *db.sql) *UserRepository {
	return &UserRepository{
		client: database.New(db),
	}
}
