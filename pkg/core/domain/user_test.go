package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/olegshishkin/financier/pkg/core/domain"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	type args struct {
		name  string
		email string
	}

	type expected struct {
		user *domain.User
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "success",
			args: args{name: "name1", email: "email1"},
			expected: expected{
				user: &domain.User{
					ID:       "",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
		},
		{
			name: "noArgs",
			args: args{name: "", email: ""},
			expected: expected{
				user: &domain.User{
					ID:       "",
					Name:     "",
					Email:    "",
					Disabled: false,
					Version:  0,
				},
			},
		},
		{
			name: "onlyName",
			args: args{name: "name1", email: ""},
			expected: expected{
				user: &domain.User{
					ID:       "",
					Name:     "name1",
					Email:    "",
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

			user := domain.NewUser(tt.args.name, tt.args.email)

			assertions.Equal(tt.expected.user, user)
		})
	}
}

func TestUser_Disable(t *testing.T) {
	t.Parallel()

	type args struct {
		user *domain.User
	}

	type expected struct {
		user    *domain.User
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "success",
			args: args{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				wantErr: false,
			},
		},
		{
			name: "noUser",
			args: args{
				user: nil,
			},
			expected: expected{
				user:    nil,
				wantErr: true,
			},
		},
		{
			name: "blankUser",
			args: args{
				user: &domain.User{
					ID:       "",
					Name:     "",
					Email:    "",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:       "",
					Name:     "",
					Email:    "",
					Disabled: false,
					Version:  0,
				},
				wantErr: true,
			},
		},
		{
			name: "userDoesNotExist",
			args: args{
				user: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				wantErr: true,
			},
		},
		{
			name: "userAlreadyDisabled",
			args: args{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				wantErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)

			err := tt.args.user.Disable()

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.user, tt.args.user)
		})
	}
}

func TestUser_Exists(t *testing.T) {
	t.Parallel()

	type args struct {
		user *domain.User
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
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: true},
		},
		{
			name: "successNotNumericId",
			args: args{
				user: &domain.User{
					ID:       "some_id",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: true},
		},
		{
			name: "noId",
			args: args{
				user: &domain.User{
					ID:       "",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "numericZeroId",
			args: args{
				user: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "numericNegativeId",
			args: args{
				user: &domain.User{
					ID:       "-3",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name: "blankUser",
			args: args{user: &domain.User{
				ID:       "",
				Name:     "",
				Email:    "",
				Disabled: false,
				Version:  0,
			}},
			expected: expected{result: false},
		},
		{
			name:     "noUser",
			args:     args{user: nil},
			expected: expected{result: false},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.user.Exists())
		})
	}
}

func TestUser_UpdateFrom(t *testing.T) {
	t.Parallel()

	type args struct {
		original    *domain.User
		mergeTarget domain.User
	}

	type expected struct {
		merged  *domain.User
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name: "success",
			args: args{
				original: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: domain.User{
					ID:       "2",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "1",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
				wantErr: false,
			},
		},
		{
			name: "noOriginal",
			args: args{
				original: nil,
				mergeTarget: domain.User{
					ID:       "1",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged:  nil,
				wantErr: true,
			},
		},
		{
			name: "noTarget",
			args: args{
				original: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  1,
				},
				mergeTarget: domain.User{
					ID:       "",
					Name:     "",
					Email:    "",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "1",
					Name:     "",
					Email:    "",
					Disabled: false,
					Version:  0,
				},
				wantErr: false,
			},
		},
		{
			name: "disabledOriginal",
			args: args{
				original: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				mergeTarget: domain.User{
					ID:       "2",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				wantErr: true,
			},
		},
		{
			name: "disabledTarget",
			args: args{
				original: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: domain.User{
					ID:       "2",
					Name:     "name2",
					Email:    "email2",
					Disabled: true,
					Version:  1,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "1",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
				wantErr: false,
			},
		},
		{
			name: "originalDoesNotExist",
			args: args{
				original: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: domain.User{
					ID:       "1",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				wantErr: true,
			},
		},
		{
			name: "targetDoesNotExist",
			args: args{
				original: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: domain.User{
					ID:       "0",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &domain.User{
					ID:       "1",
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
				wantErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)

			err := tt.args.original.UpdateFrom(tt.args.mergeTarget)

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.merged, tt.args.original)
		})
	}
}

func TestUser_String(t *testing.T) {
	t.Parallel()

	type args struct {
		user *domain.User
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
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: "name1, email1"},
		},
		{
			name: "userDoesNotExist",
			args: args{
				user: &domain.User{
					ID:       "0",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: "name1, email1"},
		},
		{
			name:     "noUser",
			args:     args{user: nil},
			expected: expected{result: ""},
		},
		{
			name: "disabledUser",
			args: args{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
			},
			expected: expected{result: "name1, email1"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.user.String())
		})
	}
}
