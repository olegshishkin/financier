package services

import "github.com/olegshishkin/financier/pkg/core/ports/output"

func GetStorage(svc *UsrSvc) *output.UserStorage {
	return &svc.storage
}

func SetStorage(svc *UsrSvc, storage output.UserStorage) {
	svc.storage = storage
}
