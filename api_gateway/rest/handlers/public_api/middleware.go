package public_api

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/public_api_errors/errors"
)

func HandlePublicApiError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		ctx := c.Request().Context()

		var errorCode uint8

		switch err.(type) {
		case *echo.HTTPError:
			return err
		case errors.AuthenticationFailed:

		}

	}
}
