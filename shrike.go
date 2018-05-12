package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{}
)
var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
	log.Level = logrus.DebugLevel
}

func main() {
	fmt.Println("Main function :")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.File("/", "static/index.html")
	e.GET("/about", about)
	e.GET("/ws", webSocketHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func about(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hi this is my cool web server.")
}

func webSocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
