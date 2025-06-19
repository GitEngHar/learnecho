package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type contextKey string

func logWithRequestID(ctx context.Context, message string) {
	if reqID, ok := ctx.Value(contextKey("request_id")).(string); ok {
		log.Printf("[reqID=%s] %s", reqID, message)
	} else {
		log.Println(message)
	}
}

func handlerTrace(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), contextKey("request_id"), "abc_123")
	defer ctx.Done()
	logWithRequestID(ctx, "StartLogTraceHandler")
	return c.String(http.StatusOK, "Logged with request ID")
}

func main() {
	e := echo.New()
	e.GET("/", handlerTrace)
	if err := e.Start(":9010"); err != nil {
		log.Fatalf(err.Error())
	}
}
