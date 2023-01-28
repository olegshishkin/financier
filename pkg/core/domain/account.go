package domain

import (
	"fmt"
	"strconv"
)

type Account struct {
	ID       string
	Name     string
	Balance  int64
	Comment  string
	Disabled bool
	Version  uint64
}

func NewAccount(name string) *Account {
	return &Account{
		ID:       "",
		Name:     name,
		Balance:  0,
		Disabled: false,
		Comment:  "",
		Version:  0,
	}
}

func (a *Account) Exists() bool {
	if a == nil || a.ID == "" {
		return false
	}

	number, err := strconv.Atoi(a.ID)
	if err != nil {
		return true
	}

	return number > 0
}

func (a *Account) String() string {
	if a == nil || (a.Name == "") {
		return ""
	}

	return fmt.Sprintf(a.Name)
}
