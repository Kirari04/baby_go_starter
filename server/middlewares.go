package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddlewares(e *echo.Echo) {
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
	}))

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Skipper: middleware.DefaultSkipper,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format:  "${time_rfc3339} ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
	}))
}
