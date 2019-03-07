package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Level stores logging level
type Level int

const (
	// OFF turn logging off
	OFF Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	ALL
)

// LogLevels maps strign to Level
var LogLevels = map[string]Level{
	"OFF":   OFF,
	"FATAL": FATAL,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"TRACE": TRACE,
	"ALL":   ALL,
}

// LogLevelLables maps Level to string
var LogLevelLables = map[Level]string{
	OFF:   "OFF",
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
	ALL:   "ALL",
}

// Logger is logger object
type Logger struct {
	sync.Mutex
	logger *log.Logger
	Level  Level
}

// InitLogger returns initialized logger
func InitLogger(filename string) (logger *Logger, err error) {
	file := os.Stdout
	if filename != "" {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("InitLogger: Failed to open log file: %v", err)
		}
	}
	return &Logger{
		logger: log.New(file, "", log.Ldate|log.Ltime),
		Level:  ALL,
	}, nil
}

func (l *Logger) Tracef(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > DEBUG {
		l.logger.SetPrefix("TRACE ")
		l.logger.Printf(format, a...)
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > INFO {
		l.logger.SetPrefix("DEBUG ")
		l.logger.Printf(format, a...)
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > WARN {
		l.logger.SetPrefix("INFO ")
		l.logger.Printf(format, a...)
	}
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > ERROR {
		l.logger.SetPrefix("WARN ")
		l.logger.Printf(format, a...)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > FATAL {
		l.logger.SetPrefix("ERROR ")
		l.logger.Printf(format, a...)
	}
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.Level > OFF {
		l.logger.SetPrefix("FATAL ")
		l.logger.Fatalf(format, a...)
	}
}
