package http

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-echo/config"
	"github.com/labstack/echo/v4"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("API-Key") != config.GetConfig().API_KEY {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
