package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/api_gateway/rest/handlers/public_api"
)

var (
	teacherRole = []string{"teacher"}
	allRoles    = []string{"teacher", "student"}
)

func registerPublicApi(e *echo.Echo) {
	api := e.Group(publicApiPrefix, public_api.HandlePublicApiError)

	auth := api.Group("/auth")
	auth.POST("/login", public_api.Login)
	auth.POST("/register", public_api.Register)

	resources := api.Group("/resources")
	resources.GET("/:resId", public_api.GetResource, public_api.AllowRoles(allRoles))
	resources.GET("/:subjId/all", public_api.GetAllResources, public_api.AllowRoles(allRoles))
	resources.POST("/:subjId", public_api.CreateResource, public_api.AllowRoles(teacherRole))
	resources.DELETE("/:resId", public_api.DeleteResource, public_api.AllowRoles(teacherRole))

	subjects := api.Group("/subjects")
	subjects.GET("/:subjId", public_api.GetSubject, public_api.AllowRoles(allRoles))
	subjects.GET("", public_api.GetAllSubjects, public_api.AllowRoles(allRoles))
	subjects.POST("", public_api.CreateSubject, public_api.AllowRoles(teacherRole))
	subjects.DELETE("/:subjId", public_api.DeleteSubject, public_api.AllowRoles(teacherRole))
}
