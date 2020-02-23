package log

import (
	"encoding/json"
	"testing"
)

func TestDEFAULT(t *testing.T) {
	out := DEFAULT("DEFAULT")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "DEFAULT" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestStackdriverMessage_StdOut(t *testing.T) {
	DEFAULT("DEFAULT").Stdout()
}

func TestDEBUG(t *testing.T) {
	out := DEBUG("DEBUG")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "DEBUG" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestINFO(t *testing.T) {
	out := INFO("INFO")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "INFO" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestNOTICE(t *testing.T) {
	out := NOTICE("NOTICE")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "NOTICE" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestWARNING(t *testing.T) {
	out := WARNING("WARNING")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "WARNING" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestERROR(t *testing.T) {
	out := ERROR("ERROR")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "ERROR" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestCRITICAL(t *testing.T) {
	out := CRITICAL("CRITICAL")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "CRITICAL" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestALERT(t *testing.T) {
	out := ALERT("ALERT")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "ALERT" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}

func TestEMERGENCY(t *testing.T) {
	out := EMERGENCY("EMERGENCY")
	log := stackdriverLog{}
	var x interface{} = out

	// stackdriverMessage type assertion
	_, ok := x.(StackdriverMessage)
	if !ok {
		t.Error("Failed. Expected type stackdriverMessage")
	}

	// Unmarshalling to check correct message data
	err := json.Unmarshal([]byte(out), &log)
	if err != nil {
		t.Error(err)
	}

	if log.Message != "EMERGENCY" {
		t.Errorf("Unmarshaliing failed expected: DEFAULT, received: %v", log.Message)
	}
}
