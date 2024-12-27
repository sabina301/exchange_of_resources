package rest

import (
	"context"
	"github.com/labstack/echo/v4"
)

const (
	serviceVersion  = "v1"
	publicApiPrefix = "/api/" + serviceVersion
	XAuthUser       = "X-Auth-User"
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
	Port string
}

func NewServer(port string) Server {
	return &ServerImpl{
		Port: port,
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
	registerPublicApi(e)
	ServerApi = e
	StartServerFunc(e, s.Port)
}
