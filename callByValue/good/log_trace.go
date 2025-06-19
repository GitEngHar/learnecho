package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	e.GET("/", handler)
	if err := e.Start(":9010"); err != nil {
		log.Fatalf(err.Error())
	}
}
