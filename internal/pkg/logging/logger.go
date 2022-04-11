package logging

import "fmt"

func LogMessage(message string) {
	fmt.Println(message)
}

func LogMessagef(format string, values ...interface{}) {
	message := fmt.Sprintf(format, values...)
	LogMessage(message)
}
