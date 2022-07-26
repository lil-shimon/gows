package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func handleServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func handleWS(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		err := websocket.Message.Send(ws, "Hello, Client")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			err := websocket.Message.Send(ws, fmt.Sprintf("\"%s\" received from client", msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handleServer)
	e.Static("/ws-client", "public")
	e.GET("/ws-client/ws", handleWS)
	e.Logger.Fatal(e.Start(":1323"))
}