package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/child", func(c echo.Context) error {
		start := time.Now()
		baseContext := c.Request().Context()
		ctx, cancel := context.WithTimeout(baseContext, 5*time.Second)
		defer cancel()
		log.Println("child received")
		select {
		case <-time.After(5 * time.Second):
			log.Println("end")
			return c.String(http.StatusOK, "done")
		case <-ctx.Done():
			end := time.Now()
			log.Println("canceled")
			log.Println(end.Sub(start))
			return ctx.Err()
		}
	})
	e.Logger.Fatal(e.Start(":9002"))
}
