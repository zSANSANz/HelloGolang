package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", HelloController)
	e.Start(":1323")
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
