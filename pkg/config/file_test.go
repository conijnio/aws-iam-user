package config

import (
	"github.com/go-ini/ini"
	"testing"
)

func TestFile(t *testing.T) {
	cfg := ini.Empty()
	section := cfg.Section("default")
	section.Key("aws_access_key_id").SetValue("OldKey")
	section.Key("aws_secret_access_key").SetValue("OldSecret")

	file := &File{}

	if err := file.writeConfig("test.ini", cfg); err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}

	loadedConfig, err := file.readConfig("test.ini")

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}

	section = loadedConfig.Section("default")

	if section.Key("aws_access_key_id").Value() != "OldKey" {
		t.Errorf("No error was expected but received: %s", err)
	}

	if section.Key("aws_secret_access_key").Value() != "OldSecret" {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestFileLoadError(t *testing.T) {
	file := &File{}

	_, err := file.readConfig("unknown-file.ini")

	if err == nil {
		t.Errorf("Error was expected but received: %s", err)
	}
}
