package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	apiGroup := e.Group("/api")
	apiGroup.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Failed to start server: ", err)
	}
}
