package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json: "name" form: "name"`
	Email string `json: "email" form: "email"`
}

func CreateUser(c echo.Context) error {
	// get data from value
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	// routing with query parameter
	e.POST("/users", CreateUser)

	e.Start(":7000")
}
