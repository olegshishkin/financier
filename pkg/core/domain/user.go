package domain

import (
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Disabled bool   `json:"-"`
	Version  int64  `json:"version"`
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
	return u.ID > 0
}

func (u *User) String() string {
	if u.Name == "" && u.Email == "" {
		return ""
	}

	return fmt.Sprintf("%s, %s", u.Name, u.Email)
}
