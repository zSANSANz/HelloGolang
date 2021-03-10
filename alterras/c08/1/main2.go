package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{Name: "Sandi", Email: "sandi@hasantech.id"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	// crate a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/user", GetUser)

	// Start the server, and log if it fails
	e.Start(":7000")
}
