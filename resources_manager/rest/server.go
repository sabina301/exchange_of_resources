package rest

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/resources_manager/repo"
)

const (
	serviceVersion    = "v1"
	internalApiPrefix = "/int/" + serviceVersion
	XAuthUser         = "X-Auth-User"
)

var (
	ServerApi       *echo.Echo
	StartServerFunc = startServer
)

type Server interface {
	Start()
}

func startServer(e *echo.Echo, port string) {
	e.Logger.Fatal(e.Start("127.0.0.1:" + port))
}

type ServerImpl struct {
	Port               string
	Repo               repo.ResourceRepository
	ResourceController *ResourceController
}

func NewServer(port string, repo repo.ResourceRepository, controller *ResourceController) Server {
	return &ServerImpl{
		Port:               port,
		Repo:               repo,
		ResourceController: controller,
	}
}

var ContextMiddleware = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetRequest(c.Request().WithContext(GetContextWithParameters(c.Request().Context(), c)))
		return next(c)
	}
}

func GetRequestId(c echo.Context) string {
	return c.Request().Header.Get(echo.HeaderXRequestID)
}

func GetAuthUser(c echo.Context) string {
	return c.Request().Header.Get(XAuthUser)
}

func GetContextWithParameters(parent context.Context, c echo.Context) context.Context {
	return context.WithValue(
		context.WithValue(
			context.WithValue(parent, echo.HeaderXRequestID, GetRequestId(c)),
			XAuthUser, GetAuthUser(c)),
		"path", c.Path())
}

func (s *ServerImpl) Start() {
	e := echo.New()
	e.Use(ContextMiddleware)
	s.registerPublicApi(e)
	ServerApi = e
	StartServerFunc(e, s.Port)
}
