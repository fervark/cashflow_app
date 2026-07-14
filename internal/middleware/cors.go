package cors

import (
	"cashflow/config"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Origins() echo.MiddlewareFunc {
	conf := config.New()

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{conf.Cors.Origin},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}
