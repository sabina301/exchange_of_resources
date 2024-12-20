package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/api_gateway/rest/handlers/public_api"
)

func registerPublicApi(e *echo.Echo) {
	apiBase := e.Group(publicApiPrefix)
	api := apiBase.Group("", public_api.HandlePublicApiError)
}
