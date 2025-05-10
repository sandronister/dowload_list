package web

import "github.com/labstack/echo/v4"

func (s *EchoServer) Start() error {
	if err := s.router.Start(s.webPort); err != nil {
		return err
	}
	return nil
}

func (s *EchoServer) AddRoute(method, path string, handlerFunc echo.HandlerFunc) {
	s.router.Add(method, path, handlerFunc)
}
