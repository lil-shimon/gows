package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func handleServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handleServer)
	e.Logger.Fatal(e.Start(":1323"))
}