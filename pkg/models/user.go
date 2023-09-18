package models

import "fmt"

type User struct {
	Account string
	UserId  string
	Arn     string
	Type    string
	Name    string
}

func (u *User) IsUser() bool {
	return u.Type == "user"
}

func (u *User) IsRole() bool {
	return u.Type == "assumed-role"
}

func (u *User) String() string {
	return fmt.Sprintf("%s in %s", u.Name, u.Account)
}
