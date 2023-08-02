package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/starcharles/sqlc-example/presentation/response"
)

type UsersController struct {
}

func (r UsersController) Get(ctx echo.Context) {
	users, err := client.GetUsers(context.Background(), 10)
	if err != nil {
		return err
	}

	var res response.GetUsersResponse
	res.Users = make([]*response.User, 0)
	for _, user := range users {
		res.Users = append(res.Users, &response.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}

	return ctx.JSON(http.StatusOK, res)

}
