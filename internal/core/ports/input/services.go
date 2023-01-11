package input

import "github.com/olegshishkin/financier/internal/core/domain"

type UserServicePort interface {
	Create(name, email string) (*domain.User, error)
	Get(email string) (*domain.User, error)
	Update(user *domain.User) error
	Disable(id uint64) error
}
