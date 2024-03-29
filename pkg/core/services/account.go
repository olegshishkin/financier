package services

import (
	"github.com/pkg/errors"

	"github.com/olegshishkin/financier/pkg/core/domain"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
)

type AccountService struct {
	storage output.AccountStorage
}

func NewAccountService(storage output.AccountStorage) *AccountService {
	return &AccountService{storage}
}

func (s *AccountService) Create(name, comment string) (*domain.Account, error) {
	if name == "" {
		return nil, errors.Errorf("invalid args, name: %s", name)
	}

	account, err := s.storage.FindEnabledByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "account not found")
	}

	if account.Exists() {
		return nil, errors.Errorf("account %s already exists", name)
	}

	account = domain.NewAccount(name, comment)

	if err = s.storage.Create(account); err != nil {
		return nil, errors.Wrap(err, "account hasn't been created")
	}

	return account, nil
}

func (s *AccountService) GetAll() ([]*domain.Account, error) {
	accounts, err := s.storage.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "can't get all accounts")
	}

	return accounts, nil
}
