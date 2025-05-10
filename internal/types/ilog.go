package types

type ILog interface {
	Info(className string, message string)
	Error(className string, message string)
}
