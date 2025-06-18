package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/child", func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Println("child server start")

		select {
		case <-time.After(5 * time.Second):
			log.Println("end")
			return c.String(http.StatusOK, "done")
		case <-ctx.Done():
			log.Println("canceled")
			return ctx.Err()
		}
	})
	e.Logger.Fatal(e.Start(":9002"))
}
