package web

func (s *EchoServer) Start() error {
	if err := s.router.Start(s.webPort); err != nil {
		return err
	}
	return nil
}
