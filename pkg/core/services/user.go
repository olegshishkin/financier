package user

import (
	"github.com/olegshishkin/financier/pkg/core/domain"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
	"strings"

	"github.com/pkg/errors"
)

type service struct {
	storage output.UserStorage
}

func NewService(storage output.UserStorage) *service {
	return &service{storage}
}

func (s *service) Create(name, email string) (user *domain.User, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "user hasn't been created")
		}
	}()

	normEmail := strings.ToLower(email)

	user, err = s.storage.FindActiveByEmail(normEmail)
	if err != nil {
		return nil, err
	}

	if user != nil {
		err = errors.Errorf("user with email %s already exists", email)
		return nil, err
	}

	user = domain.NewUser(name, normEmail)

	if err = s.storage.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Get(email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.Errorf("no email")
	}

	user, err := s.storage.FindActiveByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "user search has been failed")
	}

	return user, nil
}

func (s *service) Update(user *domain.User) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "user hasn't been updated")
		}
	}()

	if user == nil {
		err = errors.Errorf("no argument")
		return
	}

	original, err := s.storage.Get(user.ID)
	if err != nil {
		return err
	}

	if original == nil {
		return errors.Errorf("no user found")
	}

	if err = original.UpdateFrom(*user); err != nil {
		return err
	}

	if err = s.storage.Update(original); err != nil {
		return err
	}

	*user = *original

	return nil
}

func (s *service) Disable(id uint64) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "user hasn't been disabled")
		}
	}()

	user, err := s.storage.Get(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.Errorf("no user found")
	}

	if err = user.Disable(); err != nil {
		return err
	}

	if err = s.storage.Update(user); err != nil {
		return err
	}

	return nil
}
