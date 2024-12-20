package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/subjects_manager/rest/public_api"
)

func registerPublicApi(e *echo.Echo) {
	api := e.Group(publicApiPrefix)
	resources := api.Group("/resources")
	resources.GET("/:resId", public_api.GetResource)
	resources.GET("/:subjId", public_api.GetAllResources)
	resources.POST("/:subjId", public_api.CreateResource)
	resources.DELETE("/:resId", public_api.DeleteResource)
}
