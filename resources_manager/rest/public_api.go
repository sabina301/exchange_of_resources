package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/auth/rest/public_api"
)

func registerPublicApi(e *echo.Echo) {
	api := e.Group(publicApiPrefix)
	auth := api.Group("/auth")
	auth.POST("/login", public_api.Login)
	auth.POST("/register", public_api.Register)
}
