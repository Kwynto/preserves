package logger

import (
	"fmt"
	"time"
)

// const (
// 	logInfo    = "INFO"
// 	logWarning = "WARNING"
// 	logError   = "ERROR"
// )

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			// fmt.Println("End logger function.")
			break
		}
	}
}

func init() {
	go logger()
	// logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	// logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	// time.Sleep(3 * time.Second)
	// doneCh <- struct{}{}
}
