package simlog

import (
	"fmt"
	"log"
	"os"
)

const (
	LOG_M_ALL     = "GENERAL: "
	LOG_M_INFO    = "INFO: "
	LOG_M_WARNING = "WARNING: "
	LOG_M_ERR     = "ERROR: "
	LOG_M_DEBUG   = "DEBUG: "
	LOG_M_CRIT    = "CRITICAL: "

	LOG_F_PATH = "./logs/"

	LOG_F_ALL      = "all.log"
	LOG_F_INFO     = "info.log"
	LOG_F_WARNING  = "warn.log"
	LOG_F_ERR      = "err.log"
	LOG_F_DEBUG    = "debug.log"
	LOG_F_CRITICAL = "crit.log"
)

type logDescFilesS struct {
	allF,
	infoF,
	warningF,
	errF,
	debugF,
	criticalF *os.File
}

type Logger struct {
	logAll      *log.Logger
	logInfo     *log.Logger
	logWarning  *log.Logger
	logDebug    *log.Logger
	logError    *log.Logger
	logCritical *log.Logger
}

var descFiles logDescFilesS = logDescFilesS{
	allF:      &os.File{},
	infoF:     &os.File{},
	warningF:  &os.File{},
	errF:      &os.File{},
	debugF:    &os.File{},
	criticalF: &os.File{},
}

var defaultLog *Logger = new()
var isDebug bool = true

// var programName string = ""

func getDescFile(name string) (*os.File, error) {
	nLogFile := fmt.Sprintf("%s%s", LOG_F_PATH, name)
	desk, err := os.OpenFile(nLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return desk, nil
}

func new() *Logger {
	flag := log.LstdFlags

	errDir := os.MkdirAll(LOG_F_PATH, 0775)
	if errDir != nil {
		log.Fatalln("Problem to create './logs/'.")
	}

	allF, err := getDescFile(LOG_F_ALL)
	if err != nil {
		return nil
	}

	infoF, err := getDescFile(LOG_F_INFO)
	if err != nil {
		return nil
	}

	warningF, err := getDescFile(LOG_F_WARNING)
	if err != nil {
		return nil
	}

	errF, err := getDescFile(LOG_F_ERR)
	if err != nil {
		return nil
	}

	debugF, err := getDescFile(LOG_F_DEBUG)
	if err != nil {
		return nil
	}

	criticalF, err := getDescFile(LOG_F_CRITICAL)
	if err != nil {
		return nil
	}

	descFiles.allF = allF
	descFiles.infoF = infoF
	descFiles.warningF = warningF
	descFiles.errF = errF
	descFiles.debugF = debugF
	descFiles.criticalF = criticalF

	logAll := log.New(descFiles.allF, LOG_M_ALL, flag)
	logInfo := log.New(descFiles.infoF, LOG_M_INFO, flag)
	logWarning := log.New(descFiles.warningF, LOG_M_WARNING, flag)
	logDebug := log.New(descFiles.debugF, LOG_M_DEBUG, flag)
	logError := log.New(descFiles.errF, LOG_M_ERR, flag)
	logCritical := log.New(descFiles.criticalF, LOG_M_CRIT, flag)

	log.SetFlags(flag)
	log.SetPrefix(LOG_M_ALL)

	return &Logger{
		logAll:      logAll,
		logInfo:     logInfo,
		logWarning:  logWarning,
		logDebug:    logDebug,
		logError:    logError,
		logCritical: logCritical,
	}
}

func IsDebug() bool {
	return isDebug
}

func SetDebug(b bool) {
	isDebug = b
}

func EnableDebug() {
	SetDebug(true)
}

func DisableDebug() {
	SetDebug(false)
}

func GetStdLogger() *log.Logger {
	return defaultLog.logAll
}

func SetFlags(flag int) {
	if flag <= 0 {
		flag = log.LstdFlags
	}

	log.SetFlags(flag)

	defaultLog.logAll.SetFlags(flag)
	defaultLog.logInfo.SetFlags(flag)
	defaultLog.logWarning.SetFlags(flag)
	defaultLog.logDebug.SetFlags(flag)
	defaultLog.logError.SetFlags(flag)
	defaultLog.logCritical.SetFlags(flag)
}

// Print functions

func Printf(format string, v ...any) {
	defaultLog.logAll.Printf(format, v...)
	log.Printf(format, v...)
}

func Print(v ...any) {
	defaultLog.logAll.Print(v...)
	log.Print(v...)
}

func Println(v ...any) {
	defaultLog.logAll.Println(v...)
	log.Println(v...)
}

func Fatal(v ...any) {
	defaultLog.logAll.Print(v...)
	log.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	defaultLog.logAll.Printf(format, v...)
	log.Fatalf(format, v...)
}

func Fatalln(v ...any) {
	defaultLog.logAll.Println(v...)
	log.Fatalln(v...)
}

func Panic(v ...any) {
	defaultLog.logAll.Print(v...)
	log.Panic(v...)
}

func Panicf(format string, v ...any) {
	defaultLog.logAll.Printf(format, v...)
	log.Panicf(format, v...)
}

func Panicln(v ...any) {
	defaultLog.logAll.Println(v...)
	log.Panicln(v...)
}

// Level Print Function

func globPrint(msg string, prefix string) {
	defaultLog.logAll.SetPrefix(prefix)
	defaultLog.logAll.Print(msg)
	defaultLog.logAll.SetPrefix(LOG_M_ALL)

	log.SetPrefix(prefix)
	log.Print(msg)
	log.SetPrefix(LOG_M_ALL)
}

func Info(v ...any) {
	msg := fmt.Sprintln(v...)
	defaultLog.logInfo.Print(msg)

	globPrint(msg, LOG_M_INFO)
}

func Infof(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	defaultLog.logInfo.Print(msg)

	globPrint(msg, LOG_M_INFO)
}

func Warning(v ...any) {
	msg := fmt.Sprintln(v...)
	defaultLog.logWarning.Print(msg)

	globPrint(msg, LOG_M_WARNING)
}

func Warningf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	defaultLog.logWarning.Print(msg)

	globPrint(msg, LOG_M_WARNING)
}

func Debug(v ...any) {
	if isDebug {
		msg := fmt.Sprintln(v...)
		defaultLog.logDebug.Print(msg)

		globPrint(msg, LOG_M_DEBUG)
	}
}

func Debugf(format string, v ...any) {
	if isDebug {
		msg := fmt.Sprintf(format, v...)
		defaultLog.logDebug.Print(msg)

		globPrint(msg, LOG_M_DEBUG)
	}
}

func Error(v ...any) {
	msg := fmt.Sprintln(v...)
	defaultLog.logError.Print(msg)

	globPrint(msg, LOG_M_ERR)
}

func Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	defaultLog.logError.Print(msg)

	globPrint(msg, LOG_M_ERR)
}

func Critical(v ...any) {
	msg := fmt.Sprintln(v...)
	defaultLog.logCritical.Print(msg)

	globPrint(msg, LOG_M_CRIT)
}

func Criticalf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	defaultLog.logCritical.Print(msg)

	globPrint(msg, LOG_M_CRIT)
}
