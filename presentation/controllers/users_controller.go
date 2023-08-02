package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	application "github.com/starcharles/sqlc-example/application/user"
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

	application.UserFindUseCase{
		userRepository: repositories.NewUserRepository(db),
	}.Handle(application.UserFindInput{
		Id: req.Id,
	})

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
