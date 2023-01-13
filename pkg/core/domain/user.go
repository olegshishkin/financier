package domain

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Disabled bool
	Version  uint64
}

func NewUser(name, email string) *User {
	return &User{Name: name, Email: email, Disabled: false}
}

func (u *User) UpdateFrom(user User) error {
	if !u.Exists() {
		return errors.Errorf("user doesn't exist")
	}

	if u.Disabled {
		return errors.Errorf("user is disabled")
	}

	u.Name = user.Name
	u.Email = user.Email
	u.Version = user.Version

	return nil
}

func (u *User) Disable() error {
	if !u.Exists() {
		return errors.Errorf("user doesn't exist")
	}

	if u.Disabled {
		return errors.Errorf("user is already disabled")
	}

	u.Disabled = true

	return nil
}

func (u *User) Exists() bool {
	if u == nil || u.ID == "" {
		return false
	}

	number, err := strconv.Atoi(u.ID)
	if err != nil {
		return true
	}

	return number > 0
}

func (u *User) String() string {
	if u == nil || (u.Name == "" && u.Email == "") {
		return ""
	}

	return fmt.Sprintf("%s, %s", u.Name, u.Email)
}
