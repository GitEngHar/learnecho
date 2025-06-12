package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// queryParam
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// form
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

func main() {
	e := echo.New()
	skipper := func(c echo.Context) bool {
		// Skip health check endpoint
		return c.Request().URL.Path == "/health"
	}
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		Skipper:   skipper,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))
	//e.Use(middleware.Logger())
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method}, uri=${uri}, status=${status}\n",
	//}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	// http://localhost:1323/users/har
	// The return string is 'har'
	e.GET("/users/:id", getUser)

	// http://localhost:1323/show?team=dev&member=har
	// The return string is 'team:dev, member:har'
	e.GET("/show", show)

	// curl -d "name=git eng" -d "email=hoge@gmail.com" http://localhost:1323/save
	// namegit eng, email:hoge@gmail.com
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1323"))
}
