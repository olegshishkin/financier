package services

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/olegshishkin/financier/pkg/core/domain"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
)

type UsrSvc struct {
	storage output.UserStorage
}

func NewService(storage output.UserStorage) *UsrSvc {
	return &UsrSvc{storage}
}

func (s *UsrSvc) Create(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, errors.Errorf("invalid args, name: %s, email: %s", name, email)
	}

	user, err := s.storage.FindEnabledByEmail(strings.ToLower(email))
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	if user != nil {
		return nil, errors.Errorf("user with email %s already exists", email)
	}

	user = domain.NewUser(name, strings.ToLower(email))

	if err = s.storage.Create(user); err != nil {
		return nil, errors.Wrap(err, "user hasn't been created")
	}

	return user, nil
}

func (s *UsrSvc) Get(email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.Errorf("no email")
	}

	user, err := s.storage.FindEnabledByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "user search has been failed")
	}

	return user, nil
}

func (s *UsrSvc) Update(user *domain.User) error {
	if user == nil {
		return errors.Errorf("invalid arg: %s", user)
	}

	original, err := s.storage.Get(user.ID)
	if err != nil {
		return errors.Wrap(err, "getting user failed")
	}

	if original == nil {
		return errors.Errorf("no user found")
	}

	if err = original.UpdateFrom(*user); err != nil {
		return errors.Wrap(err, "user hasn't been updated")
	}

	if err = s.storage.Update(original); err != nil {
		return errors.Wrap(err, "user hasn't been updated")
	}

	*user = *original

	return nil
}

func (s *UsrSvc) Disable(id string) error {
	if id == "" {
		return errors.Errorf("invalid arg: %s", id)
	}

	user, err := s.storage.Get(id)
	if err != nil {
		return errors.Wrap(err, "getting user failed")
	}

	if user == nil {
		return errors.Errorf("no user found")
	}

	if err = user.Disable(); err != nil {
		return errors.Wrap(err, "user hasn't been disabled")
	}

	if err = s.storage.Update(user); err != nil {
		return errors.Wrap(err, "user hasn't been updated")
	}

	return nil
}
