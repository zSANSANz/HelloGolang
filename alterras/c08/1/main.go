package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/", HelloController)

	// start the server, and log if it fails
	e.Start(":7000")
}

// handler - Simple handler to make sure environtment is setup
func HelloController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status

	return c.String(http.StatusOK, "<h1>ini header</h1>")
}
