package cmd

import (
	"bytes"
	"errors"
	"github.com/conijnio/aws-iam-user/pkg/adapters"
	"github.com/conijnio/aws-iam-user/pkg/models"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

type MockAdapter struct {
	LoadUserFailure bool
}

func (u *MockAdapter) LoadUser() (*models.User, error) {
	if u.LoadUserFailure {
		return &models.User{}, errors.New("failed")
	}
	return &models.User{}, nil
}

func (u *MockAdapter) RotateCredentials(user *models.User) error {
	return nil
}

func TestRootCmd(t *testing.T) {
	var args []string

	adapters.RegisterAdapter("eu-west-1", "default", &MockAdapter{})

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRootCmdLoadUserFailure(t *testing.T) {
	var args []string

	adapters.RegisterAdapter("eu-west-1", "default", &MockAdapter{
		LoadUserFailure: true,
	})

	_, err := execute(t, rootCmd, args...)

	if err == nil {
		t.Errorf("Error was expected but received: nil")
	}
}

func TestExecute(t *testing.T) {
	var got int

	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()
	osExit = func(code int) { got = code }

	adapters.RegisterAdapter("eu-west-1", "default", &MockAdapter{})

	Execute()
	if exp := 0; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}

func TestExecuteError(t *testing.T) {
	var got int

	// Overwrite osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()
	osExit = func(code int) {
		got = code
	}

	// Overwrite rootCmd
	oldRootCmd := rootCmd
	defer func() { rootCmd = oldRootCmd }()
	rootCmd = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("mock failure")
		},
	}

	Execute()
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}

func TestRootCmdVersionNotSet(t *testing.T) {
	args := []string{"--version"}

	output, err := execute(t, rootCmd, args...)

	if err == nil {
		t.Errorf("Error was expected but received: %s", err)
	}

	if !strings.Contains(output, "unknown flag: --version") {
		t.Errorf("Expected to find 'unknown flag: --version' in: %s", output)
	}
}

func TestRootCmdVersion(t *testing.T) {
	SetVersion("0.1.0")
	args := []string{"--version"}

	output, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}

	if output != "aws-iam-user version 0.1.0" {
		t.Errorf("Expected 'aws-iam-user version 0.1.0' but received %s", output)
	}
}

func TestRootCmdDebugFlag(t *testing.T) {
	args := []string{"--debug"}

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRootCmdUnknownFlag(t *testing.T) {
	args := []string{"--unknown"}

	_, err := execute(t, rootCmd, args...)

	if err == nil {
		t.Errorf("Error was expected but received: %s", err)
	}
}
