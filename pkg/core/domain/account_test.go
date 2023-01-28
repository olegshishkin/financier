package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/olegshishkin/financier/pkg/core/domain"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
	}

	type expected struct {
		account *domain.Account
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "success",
			args: args{name: "name1"},
			expected: expected{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "",
					Disabled: false,
					Version:  0,
				},
			},
		},
		{
			name: "noArgs",
			args: args{name: ""},
			expected: expected{
				account: &domain.Account{
					ID:       "",
					Name:     "",
					Balance:  0,
					Comment:  "",
					Disabled: false,
					Version:  0,
				},
			},
		},
		{
			name: "onlyName",
			args: args{name: "name1"},
			expected: expected{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "",
					Disabled: false,
					Version:  0,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)

			account := domain.NewAccount(tt.args.name)

			assertions.Equal(tt.expected.account, account)
		})
	}
}

func TestAccount_Exists(t *testing.T) {
	t.Parallel()

	type args struct {
		account *domain.Account
	}

	type expected struct {
		result bool
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "successNumericId",
			args: args{
				account: &domain.Account{
					ID:       "1",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: true},
		},
		{
			name: "successNotNumericId",
			args: args{
				account: &domain.Account{
					ID:       "some_id",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: true},
		},
		{
			name: "noId",
			args: args{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "numericZeroId",
			args: args{
				account: &domain.Account{
					ID:       "0",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "numericNegativeId",
			args: args{
				account: &domain.Account{
					ID:       "-3",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "blankAccount",
			args: args{
				account: &domain.Account{
					ID:       "",
					Name:     "",
					Balance:  0,
					Comment:  "",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name:     "noAccount",
			args:     args{account: nil},
			expected: expected{result: false},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.account.Exists())
		})
	}
}

func TestAccount_String(t *testing.T) {
	t.Parallel()

	type args struct {
		account *domain.Account
	}

	type expected struct {
		result string
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "success",
			args: args{
				account: &domain.Account{
					ID:       "1",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: "name1"},
		},
		{
			name: "accountDoesNotExist",
			args: args{
				account: &domain.Account{
					ID:       "",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: "name1"},
		},
		{
			name:     "noAccount",
			args:     args{account: nil},
			expected: expected{result: ""},
		},
		{
			name: "disabledAccount",
			args: args{
				account: &domain.Account{
					ID:       "1",
					Name:     "name1",
					Balance:  0,
					Comment:  "comment1",
					Disabled: true,
					Version:  0,
				},
			},
			expected: expected{result: "name1"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.account.String())
		})
	}
}
