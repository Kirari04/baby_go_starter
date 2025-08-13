package server

import (
	"baby_starter/handler"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	e.GET("/", handler.GetIndex)
	e.POST("/api/user", handler.PostUser)
}
