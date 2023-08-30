package cmd

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDefaultLogging(t *testing.T) {
	if Debug != false {
		t.Errorf("Expected false but received: %t", Debug)
	}

	if log.GetLevel() != log.InfoLevel {
		t.Errorf("Expected %s but received: %s", log.InfoLevel, log.GetLevel())
	}
}

func TestToggleWithFalseLogging(t *testing.T) {
	Debug = false
	toggleDebug()

	if log.GetLevel() != log.InfoLevel {
		t.Errorf("Expected %s but received: %s", log.InfoLevel, log.GetLevel())
	}
}

func TestToggleWithTrueLogging(t *testing.T) {
	Debug = true
	toggleDebug()

	if log.GetLevel() != log.DebugLevel {
		t.Errorf("Expected %s but received: %s", log.DebugLevel, log.GetLevel())
	}
}
