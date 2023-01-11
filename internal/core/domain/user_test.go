package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		name  string
		email string
	}

	type expected struct {
		user *User
	}

	tests := []struct {
		name string
		args
		expected
	}{
		{
			name:     "success",
			args:     args{name: "name1", email: "email1"},
			expected: expected{user: &User{Name: "name1", Email: "email1"}},
		},
		{
			name:     "noArgs",
			args:     args{},
			expected: expected{user: &User{}},
		},
		{
			name:     "onlyName",
			args:     args{name: "name1"},
			expected: expected{user: &User{Name: "name1"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			user := NewUser(tt.args.name, tt.args.email)

			assertions.Equal(tt.expected.user, user)
		})
	}
}

func TestUser_Disable(t *testing.T) {
	type args struct {
		user User
	}

	type expected struct {
		user    User
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
				user: User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				user: User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				wantErr: false,
			},
		},
		{
			name: "blankUser",
			args: args{
				user: User{},
			},
			expected: expected{
				user:    User{},
				wantErr: true,
			},
		},
		{
			name: "userDoesNotExist",
			args: args{
				user: User{
					ID:       0,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{
				user: User{
					ID:       0,
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
				user: User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
			},
			expected: expected{
				user: User{
					ID:       1,
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
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			err := tt.args.user.Disable()

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.user, tt.args.user)
		})
	}
}

func TestUser_Exists(t *testing.T) {
	type args struct {
		user User
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
			name: "success",
			args: args{
				user: User{
					ID:       1,
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
				user: User{
					ID:       0,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: false},
		},
		{
			name:     "noUser",
			args:     args{user: User{}},
			expected: expected{result: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.user.Exists())
		})
	}
}

func TestUser_UpdateFrom(t *testing.T) {
	type args struct {
		original    *User
		mergeTarget User
	}

	type expected struct {
		merged  *User
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
				original: &User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: User{
					ID:       2,
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &User{
					ID:       1,
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
				wantErr: false,
			},
		},
		{
			name: "noTarget",
			args: args{
				original: &User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  1,
				},
				mergeTarget: User{},
			},
			expected: expected{
				merged: &User{
					ID:       1,
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
				original: &User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				mergeTarget: User{
					ID:       2,
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &User{
					ID:       1,
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
				original: &User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: User{
					ID:       2,
					Name:     "name2",
					Email:    "email2",
					Disabled: true,
					Version:  1,
				},
			},
			expected: expected{
				merged: &User{
					ID:       1,
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
				original: &User{
					ID:       0,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: User{
					ID:       2,
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &User{
					ID:       0,
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
				original: &User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				mergeTarget: User{
					ID:       0,
					Name:     "name2",
					Email:    "email2",
					Disabled: false,
					Version:  1,
				},
			},
			expected: expected{
				merged: &User{
					ID:       1,
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
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			err := tt.args.original.UpdateFrom(tt.args.mergeTarget)

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(*tt.expected.merged, *tt.args.original)
		})
	}
}

func TestUser_String(t *testing.T) {
	type args struct {
		user User
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
				user: User{
					ID:       1,
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
			},
			expected: expected{result: "name1, email1"},
		},
		{
			name: "noId",
			args: args{
				user: User{
					ID:       0,
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
			args:     args{user: User{}},
			expected: expected{result: ""},
		},
		{
			name: "disabledUser",
			args: args{
				user: User{
					ID:       1,
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
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			assertions.Equal(tt.expected.result, tt.args.user.String())
		})
	}
}
