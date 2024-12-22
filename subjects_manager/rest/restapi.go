package rest

import (
	"github.com/labstack/echo/v4"
)

func (s *ServerImpl) registerPublicApi(e *echo.Echo) {
	api := e.Group(internalApiPrefix)
	subjects := api.Group("/subjects")
	subjects.GET("/:subjId", s.SubjectController.GetSubject)
	subjects.GET("", s.SubjectController.GetAllSubjects)
	subjects.POST("", s.SubjectController.CreateSubject)
	subjects.DELETE("/:subjId", s.SubjectController.DeleteSubject)
}
