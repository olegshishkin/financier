package output

import "github.com/olegshishkin/financier/pkg/core/domain"

type UserStorage interface {
	Create(user *domain.User) error
	Get(id string) (*domain.User, error)
	FindEnabledByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
}

type AccountStorage interface {
	Create(account *domain.Account) error
	FindEnabledByName(name string) (*domain.Account, error)
	FindAll() ([]*domain.Account, error)
}
