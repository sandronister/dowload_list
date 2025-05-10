package web

import "github.com/labstack/echo/v4"

type EchoServer struct {
	router  *echo.Echo
	webPort string
}

func NewServer(webPort string) *EchoServer {
	return &EchoServer{
		router:  echo.New(),
		webPort: webPort,
	}
}
