package input

import "github.com/olegshishkin/financier/pkg/core/domain"

type UserService interface {
	Create(name, email string) (*domain.User, error)
	Get(email string) (*domain.User, error)
	Update(user *domain.User) error
	Disable(id string) error
}
