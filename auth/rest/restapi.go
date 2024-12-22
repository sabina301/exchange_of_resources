package rest

import (
	"github.com/labstack/echo/v4"
)

func (s *ServerImpl) registerPublicApi(e *echo.Echo) {
	api := e.Group(internalApiPrefix)
	auth := api.Group("/auth")
	auth.POST("/login", s.UserController.Login)
	auth.POST("/register", s.UserController.Register)
}
