package output

import "github.com/olegshishkin/financier/pkg/core/domain"

type UserStorage interface {
	Create(user *domain.User) error
	Get(id string) (*domain.User, error)
	FindActiveByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
}
