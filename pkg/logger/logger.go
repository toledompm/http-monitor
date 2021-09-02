package logger

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (*Logger) Error(message string, err error) {
	println("[ERROR]", message, err.Error())
}

func (*Logger) Info(message string) {
	println("[INFO]", message)
}

func (*Logger) Warn(message string) {
	println("[WARN]", message)
}
