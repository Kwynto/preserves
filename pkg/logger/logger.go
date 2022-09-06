package logger

import (
	"log"
	"os"
)

type Logging interface {
	Info(err error)
	Warning(err error)
	Debug(err error)
	Error(err error)
	Fatal(err error)
	Trace(err error)
	All(err error)
}

type logging struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	debugLogger   *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	traceLogger   *log.Logger
	allLogger     *log.Logger
}

func New(outPatch string) Logging {

	return &logging{
		infoLogger:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stderr, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger:   log.New(os.Stderr, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime),
		fatalLogger:   log.New(os.Stderr, "FATAL\t", log.Ldate|log.Ltime),
		traceLogger:   log.New(os.Stderr, "TRACE\t", log.Ldate|log.Ltime),
		allLogger:     log.New(os.Stderr, "ALL\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *logging) Info(err error) {
	//
}

func (l *logging) Warning(err error) {
	//
}

func (l *logging) Debug(err error) {
	//
}

func (l *logging) Error(err error) {
	//
}

func (l *logging) Fatal(err error) {
	//
}

func (l *logging) Trace(err error) {
	//
}

func (l *logging) All(err error) {
	//
}
