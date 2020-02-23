package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type stackdriverLog struct {
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Status    uint   `json:"status"`
}

type StackdriverMessage string

// Stdout print to STDOUT
func (r StackdriverMessage) Stdout() {
	_, err := os.Stdout.WriteString(string(r))
	if err != nil {
		fmt.Println(err.Error())
	}
}

// formatLog format json
func formatLog(v interface{}) StackdriverMessage {
	output, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}
	return StackdriverMessage(fmt.Sprintln(string(output)))
}

//DEFAULT write log message to stdout
func DEFAULT(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "DEFAULT",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    0,
	}
	return formatLog(log)
}

//DEBUG write log message to stdout
func DEBUG(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "DEBUG",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    100,
	}
	return formatLog(log)
}

//INFO write log message to stdout
func INFO(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "INFO",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    200,
	}
	return formatLog(log)
}

//NOTICE write log message to stdout
func NOTICE(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "NOTICE",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    300,
	}
	return formatLog(log)
}

//WARNING write log message to stdout
func WARNING(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "WARNING",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    400,
	}
	return formatLog(log)
}

//ERROR write log message to stdout
func ERROR(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "ERROR",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    500,
	}
	return formatLog(log)
}

//CRITICAL write log message to stdout
func CRITICAL(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "CRITICAL",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    600,
	}
	return formatLog(log)
}

//ALERT write log message to stdout
func ALERT(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "ALERT",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    700,
	}
	return formatLog(log)
}

//EMERGENCY write log message to stdout
func EMERGENCY(m string) StackdriverMessage {
	log := stackdriverLog{
		Severity:  "EMERGENCY",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    800,
	}
	return formatLog(log)
}
