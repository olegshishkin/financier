package stub

import (
	"github.com/pkg/errors"

	"github.com/olegshishkin/financier/pkg/core/domain"
)

type AccountStorageStub struct {
	accounts []*domain.Account
}

func NewAccountStorageStub() *AccountStorageStub {
	return &AccountStorageStub{
		accounts: []*domain.Account{},
	}
}

func (s *AccountStorageStub) Create(account *domain.Account) error {
	if account == nil {
		return errors.Errorf("nil passed as account")
	}

	if account.Exists() {
		return errors.Errorf("account %s already exists", account)
	}

	account.ID = "1"

	return nil
}

func (s *AccountStorageStub) FindEnabledByName(name string) (*domain.Account, error) {
	if name == "" {
		return nil, errors.Errorf("no account name")
	}

	for _, a := range s.accounts {
		if a.Name == name {
			return a, nil
		}
	}

	return nil, nil
}
