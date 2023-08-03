package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/starcharles/sqlc-example/domain/entities"
	domain "github.com/starcharles/sqlc-example/domain/repositories"
	"github.com/starcharles/sqlc-example/domain/value_object"
	"github.com/starcharles/sqlc-example/sqlc"
)

type UserRepository struct {
	client *sqlc.Queries
}

func (r UserRepository) FindOne(ctx context.Context, id value_object.ID) (*entities.User, error) {
	fmt.Println("aaaaaaa")
	fmt.Println(id.Value())
	user, err := r.client.GetUser(ctx, int64(id.Value()))
	fmt.Print(user)
	fmt.Print(err)
	if err != nil {
		return nil, err
	}

	return modelToEntity(&user)
}

func modelToEntity(model *sqlc.User) (*entities.User, error) {
	id, err := value_object.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}

	return entities.NewUser(
		*id,
		model.Name,
		model.Email,
		model.Password,
		model.CreatedAt,
		model.UpdatedAt,
	), nil
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &UserRepository{
		client: sqlc.New(db),
	}
}
