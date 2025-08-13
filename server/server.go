package server

import (
	"baby_starter/app"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	initMiddlewares(e)
	initRoutes(e)

	app.LOG.Fatal().Err(e.Start(app.ENV.Addr)).Msg("failed to start server")
}
