package rest

import (
	"github.com/labstack/echo/v4"
)

const (
	serviceVersion  = "v1"
	servicePrefix   = "/gateway/" + serviceVersion
	publicApiPrefix = "/api/" + serviceVersion
)

type Server interface {
	Start()
}

type ServerImpl struct {
	Port string
}

func NewServer(port string) *ServerImpl {
	return &ServerImpl{Port: port}
}

func (s *ServerImpl) Start() {
	e := echo.New()

	registerPublicApi(e)
}
