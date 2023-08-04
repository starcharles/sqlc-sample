package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/starcharles/sqlc-example/presentation/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.GET("/users/:id", controllers.UsersController{}.Get)
	e.Logger.Fatal(e.Start(":8989"))
}
