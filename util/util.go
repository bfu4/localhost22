package util

import "log"

// InitLogger initialize a logger with the specified name
func InitLogger(name string) {
	// magenta [name]
	const fmt = "\033[1;35m[{}]\u0020\u001B[0m"
	log.SetFlags(log.Lmsgprefix + log.Ltime)
	log.SetPrefix(Format(fmt, name))
	Info("Set up logger with name [{}]!", name)
}

// Fatal log a fatal message and exit the program
func Fatal(message string, v ...string) {
	fmt := Format("\033[1;31mERROR: {}\033[0m", message)
	log.Fatal(Format(fmt, v...))
}

func Warn(message string, v ...string) {
	fmt := Format("\033[1;31mWARN: {}\033[0m", message)
	log.Println(Format(fmt, v...))
}

// Info log an informative message
func Info(message string, v ...string) {
	log.Println(Format(message, v...))
}

// Format a string
func Format(message string, v ...string) string {
	iter := 0

	for index, character := range message {
		if index != len(message) && character == '{' && message[index+1] == '}' {
			iter++
			return message[:index] + v[iter-1] + Format(message[index+2:], v[iter:]...)
		}
	}

	return message
}

func ErrorOrMessage(err error, message string) string {
	if err == nil {
		return message
	}

	return err.Error()
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
