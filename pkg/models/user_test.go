package models

import "testing"

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
