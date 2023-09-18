package config

import (
	"errors"
	"github.com/go-ini/ini"
	"testing"
)

type MockFile struct{}

func (f *MockFile) readConfig(filename string) (*ini.File, error) {
	cfg := ini.Empty()
	section := cfg.Section("default")
	section.Key("aws_access_key_id").SetValue("OldKey")
	section.Key("aws_secret_access_key").SetValue("OldSecret")
	return cfg, nil
}

func (f *MockFile) writeConfig(filename string, cfg *ini.File) error {
	return nil
}

func TestConfig(t *testing.T) {
	cfg := New("default")
	cfg.file = &MockFile{}

	err := cfg.UpdateKeys("NewKey", "NewSecret")

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}

	if cfg.AccessKeyId() != "OldKey" {
		t.Errorf("Expected the old key to be OldKey but it is: %s", cfg.AccessKeyId())
	}

	if cfg.secretAccessKeyId != "OldSecret" {
		t.Errorf("Expected the old key to be OldSecret but it is: %s", cfg.secretAccessKeyId)
	}

	err = errors.New("my error message")
	if !errors.Is(err, cfg.Rollback(err)) {
		t.Errorf("Expected the passed error to be returned: %s", err)
	}
}
