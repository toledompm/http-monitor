package logger

func Error(message string, err error) {
	println("[ERROR]", message, err.Error())
}

func Info(message string) {
	println("[INFO]", message)
}

func Warn(message string) {
	println("[WARN]", message)
}
