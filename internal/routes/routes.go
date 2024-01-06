package routes

import (
	"context"

	"github.com/RaniAgus/go-templ/internal/users"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	app := echo.New()

	userHandler := users.UserHandler{}

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(
				c.Request().Context(),
				"version",
				"v1",
			)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})
	app.GET("/users", userHandler.HandleUserShow)

	return app
}
