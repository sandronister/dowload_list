package types

import "errors"

var ErrNilBroker = errors.New("broker is nil")
var ErrNilEnviroment = errors.New("enviroment is nil")
var ErrNilLogger = errors.New("logger is nil")
var ErrNilWebServer = errors.New("web server is nil")
var ErrNilUseCase = errors.New("use case is nil")
var ErrNilService = errors.New("service is nil")
var ErrNilConfig = errors.New("config is nil")
var ErrNilRepository = errors.New("repository is nil")
var ErrNilData = errors.New("data is nil")
