package logging

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)


type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

type BasicLogger struct {
	Level int
	Prefix string
	Logger *log.Logger
}
const (
	// LEVEL_DEBUG = 0
	LEVEL_DEBUG = iota
	// LEVEL_INFO = 1
	LEVEL_INFO
	// LEVEL_WARNING = 2
	LEVEL_WARNING
	// LEVEL_ERROR = 3
	LEVEL_ERROR
	// LEVEL_CRITICAL = 4
	LEVEL_CRITICAL
)

var (
	// errinvalid is used when an invalid log level has been used
	ErrInvalidLoglevel = errors.New("invalid log level")
	defaultLogger	   = BasicLogger{
		Level:  LEVEL_CRITICAL,
		Prefix: "",
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	logLevels 			= map[string]int{
		"LEVEL_DEBUG": LEVEL_DEBUG,
		"INFO":	LEVEL_INFO,
		"WARNING": LEVEL_WARNING,
		"ERROR" : LEVEL_ERROR,
		"CRITICAL" : LEVEL_CRITICAL,
	}

)

func NewLogger(level string , out io.Writer , prefix string) (BasicLogger,error)  {
	l,ok := logLevels[strings.ToUpper(level)]
	if !ok  {
		return defaultLogger, ErrInvalidLoglevel
	}
	return BasicLogger{Level: l,Prefix: prefix,Logger: log.New(out,"",log.LstdFlags)} , nil
}



func (l BasicLogger) Info(v ...interface{}) {
	if l.Level > LEVEL_INFO {
		return
	}
	l.prependLog("INFO",v...)
}
func (l BasicLogger) Debug(v ...interface{}) {
	if l.Level > LEVEL_DEBUG {
		return
	}
	l.prependLog("DEBUG",v...)
}
// Critical logs a message using CRITICAL as log level.
func (l BasicLogger) Critical(v ...interface{}) {
	l.prependLog("CRITICAL:", v...)
}
func (l BasicLogger) Warning(v ...interface{}) {
	if l.Level > LEVEL_WARNING {
		return
	}
	l.prependLog("INFO",v...)
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l BasicLogger) Fatal(v ...interface{}) {
	l.prependLog("FATAL:", v...)
	os.Exit(1)
}




func (l BasicLogger) prependLog(level string , v ...interface{})  {
	msg := make([]interface{},len(v)+2)
	msg[0] = l.Prefix
	msg[1] = level
	copy(msg[:2],v[:])
	l.Logger.Println(msg...)
}


