package models

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	user := &User{Type: "user"}

	if user.IsRole() {
		t.Error("Expected IsRole() to return True")
	}

	if !user.IsUser() {
		t.Error("Expected IsUser() to return False")
	}
}

func TestRole(t *testing.T) {
	user := &User{Type: "assumed-role"}

	if !user.IsRole() {
		t.Error("Expected IsRole() to return False")
	}

	if user.IsUser() {
		t.Error("Expected IsUser() to return True")
	}
}

func TestString(t *testing.T) {
	user := &User{Account: "111122223333", Name: "john.doe"}
	if fmt.Sprint(user) != "john.doe in 111122223333" {
		t.Errorf("Expected 'john.doe in 111122223333' but received '%s'", user)
	}
}
