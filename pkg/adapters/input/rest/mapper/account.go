package mapper

import (
	"strconv"

	"github.com/olegshishkin/olsh-go-utils/types"
	"github.com/pkg/errors"

	"github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/pkg/core/domain"
)

func AccountToAccountOut(src *domain.Account) *v1.AccountOutput {
	if src == nil {
		return nil
	}

	return &v1.AccountOutput{
		Balance:  src.Balance,
		Comment:  &src.Comment,
		Disabled: src.Disabled,
		Id:       src.ID,
		Name:     src.Name,
		Version:  strconv.FormatUint(src.Version, 10),
	}
}

func AccountOutToAccount(src *v1.AccountOutput) (*domain.Account, error) {
	if src == nil {
		return nil, nil
	}

	ver, err := strconv.ParseUint(src.Version, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "field 'Version' hasn't been parsed")
	}

	acc := &domain.Account{
		ID:       src.Id,
		Name:     src.Name,
		Balance:  src.Balance,
		Comment:  types.PointerVal(src.Comment),
		Disabled: src.Disabled,
		Version:  ver,
	}

	return acc, nil
}
