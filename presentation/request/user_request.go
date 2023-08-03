package request

import "github.com/starcharles/sqlc-example/domain/value_object"

type GetUserRequest struct {
	Id value_object.ID `param:"id"`
}
