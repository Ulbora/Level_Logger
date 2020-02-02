package levellogger

import "log"

const (
	//InfoLevel InfoLevel
	InfoLevel = 1
	//DebugLevel DebugLevel
	DebugLevel = 2
	//AllLevel AllLevel
	AllLevel = 3
)

//Logger Logger
type Logger struct {
	LogLevel int
}

//Info Info
func (l *Logger) Info(val interface{}) {
	if l.LogLevel == InfoLevel || l.LogLevel == AllLevel {
		log.Println("INFO: ", val)
	}
}

//Debug Debug
func (l *Logger) Debug(val interface{}) {
	if l.LogLevel == DebugLevel || l.LogLevel == AllLevel {
		log.Println("DEBUG: ", val)
	}
}

//Error Error
func (l *Logger) Error(val interface{}) {
	log.Println("ERROR: ", val)
}

//go mod init github.com/Ulbora/Level_Logger
