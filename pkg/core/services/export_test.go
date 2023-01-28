package services

import "github.com/olegshishkin/financier/pkg/core/ports/output"

func GetUserStorage(svc *UsrSvc) *output.UserStorage {
	return &svc.storage
}

func SetUserStorage(svc *UsrSvc, storage output.UserStorage) {
	svc.storage = storage
}
