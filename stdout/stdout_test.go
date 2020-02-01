package stdout

import (
	"encoding/json"
	"testing"
)

func TestStdout(t *testing.T) {
	log := stackdriverLog{
		Severity:  "DEFAULT",
		Message:   "DEFAULT MESSAGE",
		Timestamp: "",
		Status:    0,
	}

	if log.Status != 0 {
		t.Errorf("Failed. Expected value: 0. Received:%d", log.Status)
	}

	out := formatLog(log)
	err := json.Valid([]byte(out))
	if err != true {
		t.Error(err)
	}
}
