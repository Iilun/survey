package log

import (
	"log"
	"os"
)

var logEnabled = false

func init() {
	if filePath, found := os.LookupEnv("SURVEY_DEBUG"); found {
		if file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755); err == nil {
			logEnabled = true
			log.SetOutput(file)
		}
	}
}

func Println(v ...interface{}) {
	if logEnabled {
		log.Println(v...)
	}
}

func Printf(format string, v ...interface{}) {
	if logEnabled {
		log.Printf(format, v...)
	}
}
