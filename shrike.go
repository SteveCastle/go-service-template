package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
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
	e.GET("/", hello)
	e.GET("/signature", signature)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world.")
}

func signature(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hi this is my cool web server.")
}
