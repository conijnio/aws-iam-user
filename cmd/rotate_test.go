package cmd

import (
	"github.com/conijnio/aws-iam-user/pkg/adapters"
	"testing"
)

func TestRotateCmd(t *testing.T) {
	args := []string{"rotate"}

	adapters.RegisterAdapter("eu-west-1", "default", &MockAdapter{})

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRotateCmdLoadUserFailure(t *testing.T) {
	args := []string{"rotate"}

	adapters.RegisterAdapter("eu-west-1", "default", &MockAdapter{
		LoadUserFailure: true,
	})

	_, err := execute(t, rootCmd, args...)

	if err == nil {
		t.Errorf("Error was expected but received: nil")
	}
}
