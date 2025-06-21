package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func businessLogic(c context.Context) error {
	if userID, ok := c.Value("user_id").(string); ok {
		fmt.Printf("userID={%s}", userID)
		return nil
	} else {
		fmt.Printf("user id none")
		return errors.New("user id none")
	}
}

func handler(c echo.Context) error {
	// key is not correct.
	ctx := context.WithValue(c.Request().Context(), "user_id", "bad_practice_name")
	// ctx send to business logic
	if err := businessLogic(ctx); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "success")
}

func main() {
	e := echo.New() // with value 値の重複
	e.GET("/", handler)
	e.Start(":9011")
}
