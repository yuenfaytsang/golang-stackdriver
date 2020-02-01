package stdout

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

func formatLog(v interface{}) string {
	output, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}
	return fmt.Sprintln(string(output))
}

//DEFAULT write logmessange to stdout
func DEFAULT(m string) {
	log := stackdriverLog{
		Severity:  "DEFAULT",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    0,
	}
	os.Stdout.WriteString(formatLog(log))
}

//DEBUG write logmessange to stdout
func DEBUG(m string) {
	log := stackdriverLog{
		Severity:  "DEBUG",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    100,
	}
	os.Stdout.WriteString(formatLog(log))
}

//INFO write logmessange to stdout
func INFO(m string) {
	log := stackdriverLog{
		Severity:  "INFO",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    200,
	}
	os.Stdout.WriteString(formatLog(log))
}

//NOTICE write logmessange to stdout
func NOTICE(m string) {
	log := stackdriverLog{
		Severity:  "NOTICE",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    300,
	}
	os.Stdout.WriteString(formatLog(log))
}

//WARNING write logmessange to stdout
func WARNING(m string) {
	log := stackdriverLog{
		Severity:  "WARNING",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    400,
	}
	os.Stdout.WriteString(formatLog(log))
}

//ERROR write logmessange to stdout
func ERROR(m string) {
	log := stackdriverLog{
		Severity:  "ERROR",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    500,
	}
	os.Stdout.WriteString(formatLog(log))
}

//CRITICAL write logmessange to stdout
func CRITICAL(m string) {
	log := stackdriverLog{
		Severity:  "CRITICAL",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    600,
	}
	os.Stdout.WriteString(formatLog(log))
}

//ALERT write logmessange to stdout
func ALERT(m string) {
	log := stackdriverLog{
		Severity:  "ALERT",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    700,
	}
	os.Stdout.WriteString(formatLog(log))
}

//EMERGENCY write logmessange to stdout
func EMERGENCY(m string) {
	log := stackdriverLog{
		Severity:  "EMERGENCY",
		Message:   m,
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    800,
	}
	os.Stdout.WriteString(formatLog(log))
}
