package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareBasicAuth(username string, password string) echo.MiddlewareFunc {
	c := middleware.DefaultBasicAuthConfig
	c.Validator = func(s string, s2 string, context echo.Context) (bool, error) {
		if username == s && password == s2 {
			return true, nil
		}
		return false, nil
	}
	return middleware.BasicAuthWithConfig(c)
}
