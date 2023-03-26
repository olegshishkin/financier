package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "github.com/olegshishkin/financier/mocks/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/domain"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/services"
)

func createAccountService(storage output.AccountStorage) *services.AccountService {
	target := &services.AccountService{}
	services.SetAccountStorage(target, storage)

	return target
}

func TestAccountService_NewAccountService(t *testing.T) {
	t.Parallel()
	storageMock := mocks.NewAccountStorage(t)
	svc := services.NewAccountService(storageMock)
	assert.Equal(t, storageMock, *services.GetAccountStorage(svc))
}

func TestAccountService_Create(t *testing.T) {
	t.Parallel()

	type findEnabledByNameArgs struct {
		name string
	}

	type findEnabledByNameReturn struct {
		account *domain.Account
		err     error
	}

	type createArgs struct {
		account *domain.Account
	}

	type createReturn struct {
		err error
	}

	type args struct {
		name    string
		comment string
	}

	type expected struct {
		account *domain.Account
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
		findEnabledByNameArgs
		findEnabledByNameReturn
		createArgs
		createReturn
	}{
		{
			name: "success",
			args: args{name: "name1", comment: "com1"},
			expected: expected{
				account: &domain.Account{
					ID:       "1",
					Name:     "name1",
					Balance:  0,
					Comment:  "com1",
					Disabled: false,
					Version:  0,
				},
				wantErr: false,
			},
			findEnabledByNameArgs:   findEnabledByNameArgs{name: "name1"},
			findEnabledByNameReturn: findEnabledByNameReturn{account: nil, err: nil},
			createArgs: createArgs{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "com1",
					Disabled: false,
					Version:  0,
				},
			},
			createReturn: createReturn{err: nil},
		},
		{
			name:     "noName",
			args:     args{name: "", comment: "com1"},
			expected: expected{account: nil, wantErr: true},
		},
		{
			name:                  "accountAlreadyExists",
			args:                  args{name: "name1", comment: "com1"},
			expected:              expected{account: nil, wantErr: true},
			findEnabledByNameArgs: findEnabledByNameArgs{name: "name1"},
			findEnabledByNameReturn: findEnabledByNameReturn{
				account: &domain.Account{
					ID:       "1",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
				err: nil,
			},
		},
		{
			name:                    "findAccountFailed",
			args:                    args{name: "name1", comment: "com1"},
			expected:                expected{account: nil, wantErr: true},
			findEnabledByNameArgs:   findEnabledByNameArgs{name: "name1"},
			findEnabledByNameReturn: findEnabledByNameReturn{account: nil, err: services.ErrExample},
		},
		{
			name:                    "createFailed",
			args:                    args{name: "name1", comment: "com1"},
			expected:                expected{account: nil, wantErr: true},
			findEnabledByNameArgs:   findEnabledByNameArgs{name: "name1"},
			findEnabledByNameReturn: findEnabledByNameReturn{account: nil, err: nil},
			createArgs: createArgs{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "com1",
					Disabled: false,
					Version:  0,
				},
			},
			createReturn: createReturn{err: services.ErrExample},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)

			storageMock := mocks.NewAccountStorage(t)
			target := createAccountService(storageMock)

			if tt.findEnabledByNameArgs.name != "" {
				storageMock.EXPECT().
					FindEnabledByName(tt.findEnabledByNameArgs.name).
					Return(tt.findEnabledByNameReturn.account, tt.findEnabledByNameReturn.err)
			}

			if tt.createArgs.account != nil {
				storageMock.EXPECT().
					Create(tt.createArgs.account).
					Run(func(account *domain.Account) {
						if !account.Exists() {
							account.ID = "1"
						}
					}).
					Return(tt.createReturn.err)
			}

			account, err := target.Create(tt.args.name, tt.args.comment)

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.account, account)
		})
	}
}
