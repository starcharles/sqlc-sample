package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/starcharles/sqlc-example/database"
	"github.com/starcharles/sqlc-example/response"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println(os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	client := database.New(db)
	e := echo.New()

	e.GET("/users", func(ctx echo.Context) error {
		users, err := client.GetUsers(context.Background(), 10)
		if err != nil {
			return err
		}

		var res response.GetUsersResponse
		for _, user := range users {
			res.Users = append(res.Users, &response.User{
				Id:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.UpdatedAt.String(),
			})
		}

		log.Printf("%+v", res)

		return ctx.JSON(http.StatusOK, res)
	})

	// e.POST("/user", func(c echo.Context) error {
	// 	if _, err := client.CreateUser(context.Background(), c.FormValue("content")); err != nil {
	// 		log.Println(err.Error())
	// 		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	// 	}
	// 	return c.Redirect(http.StatusFound, "/")
	// })
	e.Logger.Fatal(e.Start(":8989"))
}
