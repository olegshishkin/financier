package services

import (
	"errors"

	"github.com/olegshishkin/financier/pkg/core/ports/output"
)

var ErrExample = errors.New("")

func GetUserStorage(svc *UsrSvc) *output.UserStorage {
	return &svc.storage
}

func SetUserStorage(svc *UsrSvc, storage output.UserStorage) {
	svc.storage = storage
}
