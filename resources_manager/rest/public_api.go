package rest

import (
	"github.com/labstack/echo/v4"
)

func (s *ServerImpl) registerPublicApi(e *echo.Echo) {
	api := e.Group(internalApiPrefix)
	resources := api.Group("/resources")
	resources.GET("/:resId", s.ResourceController.GetResource)
	resources.GET("/:subjId/all", s.ResourceController.GetAllResources)
	resources.POST("/:subjId", s.ResourceController.CreateResource)
	resources.DELETE("/:resId", s.ResourceController.DeleteResource)
}
