package eventlog

import (
	"fmt"
	"sync"
	"time"
)

const maxEvents = 1000

type EventLog struct {
	mu   sync.Mutex
	logs []string
}

var globalEventLog = &EventLog{
	logs: make([]string, 0, maxEvents),
}

func Add(event string) {
	globalEventLog.mu.Lock()
	defer globalEventLog.mu.Unlock()

	if len(globalEventLog.logs) >= maxEvents {
		globalEventLog.logs = globalEventLog.logs[1:] // Remove oldest event
	}

	currentTime := time.Now()
	timestamp := fmt.Sprintf("[%02d:%02d:%02d:%03d] - ", currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond()/1e6)
	event = timestamp + event
	globalEventLog.logs = append(globalEventLog.logs, event)
}

func GetAll() []string {
	globalEventLog.mu.Lock()
	defer globalEventLog.mu.Unlock()

	// Return a copy of the logs to avoid external modification
	logsCopy := make([]string, len(globalEventLog.logs))
	copy(logsCopy, globalEventLog.logs)
	return logsCopy
}
