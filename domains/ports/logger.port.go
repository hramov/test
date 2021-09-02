package ports

type LoggerPort interface {
	Log(message string)
}
