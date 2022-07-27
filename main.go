package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
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

func sendDataEachSecond(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for range time.Tick(1000 * time.Millisecond) {
			err := websocket.Message.Send(ws, "noise")
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
	e.GET("/ws/noise", sendDataEachSecond)
	e.Logger.Fatal(e.Start(":1323"))
}
