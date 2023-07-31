package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/starcharles/sqlc-example/database"

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

	e.GET("/entries", func(ctx echo.Context) error {
		entries, err := client.GetEntries(context.Background(), 10)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, entries)
	})

	e.POST("/add", func(c echo.Context) error {
		if _, err := client.CreateEntry(context.Background(), c.FormValue("content")); err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
