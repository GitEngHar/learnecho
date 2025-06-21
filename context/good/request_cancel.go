package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func slowOperation(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		return "success", nil
	case <-ctx.Done():
		return "failed", ctx.Err()
	}
}

func handler(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := slowOperation(ctx)
	if err != nil {
		fmt.Println("canceled")
		return c.String(http.StatusRequestTimeout, "request canceled : "+err.Error())
	}
	return c.String(http.StatusOK, result)
}

func main() {
	e := echo.New()
	e.GET("/", handler)
	if err := e.Start(":9010"); err != nil {
		log.Fatalf(err.Error())
	}
}
