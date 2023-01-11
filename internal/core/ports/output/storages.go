package output

import "github.com/olegshishkin/financier/internal/core/domain"

type UserStoragePort interface {
	Create(user *domain.User) error
	Get(id uint64) (*domain.User, error)
	FindActiveByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
}
