package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	user_usecase "github.com/starcharles/sqlc-example/application/user"
	"github.com/starcharles/sqlc-example/di"
	"github.com/starcharles/sqlc-example/presentation/request"
	"github.com/starcharles/sqlc-example/presentation/response"
)

type UsersController struct {
}

func (r UsersController) Get(ctx echo.Context) error {
	var req request.GetUserRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	fmt.Println(req.Id)
	fmt.Println(req.Id.Value())

	handler, err := di.UserFindUseCase(user_usecase.UserFindInput{
		ID: req.Id,
	})

	if err != nil {
		return err
	}

	res, err := handler.Handle(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.GetUserResponse{
		User: response.User{
			Id:        int64(res.User.ID().Value()),
			Name:      res.User.Name(),
			Email:     res.User.Email(),
			CreatedAt: res.User.CreatedAt(),
			UpdatedAt: res.User.UpdatedAt(),
		},
	})
}
