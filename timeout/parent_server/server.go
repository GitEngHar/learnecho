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
	e.GET("/cancel", func(c echo.Context) error {
		baseContext := c.Request().Context()
		ctx, cancel := context.WithTimeout(baseContext, 2*time.Second)
		defer cancel() // open resource
		// create request for send to server2
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9002/child", nil)
		// send to server2
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("server1: error calling server2:", err)
			return err
		}
		defer resp.Body.Close()
		//_, _ := io.ReadAll(resp.Body)
		//return c.String(http.StatusOK, string(body))
		return c.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":9001"))
}
