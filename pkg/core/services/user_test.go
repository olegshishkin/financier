package user

import (
	"errors"
	mocks "github.com/olegshishkin/financier/mocks/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_NewService(t *testing.T) {
	storageMock := mocks.NewUserStorage(t)
	assert.Equal(t, storageMock, NewService(storageMock).storage)
}

func TestService_Create(t *testing.T) {
	type findActiveByEmailArgs struct {
		email string
	}

	type findActiveByEmailReturn struct {
		user *domain.User
		err  error
	}

	type createArgs struct {
		user *domain.User
	}

	type createReturn struct {
		err error
	}

	type args struct {
		name  string
		email string
	}

	type expected struct {
		user    *domain.User
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
		findActiveByEmailArgs
		findActiveByEmailReturn
		createArgs
		createReturn
	}{
		{
			name: "success",
			args: args{name: "name1", email: "email1"},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				wantErr: false,
			},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: nil},
			createArgs: createArgs{
				user: &domain.User{
					ID:      "",
					Name:    "name1",
					Email:   "email1",
					Version: 0,
				},
			},
			createReturn: createReturn{err: nil},
		},
		{
			name: "upperCasedEmail",
			args: args{name: "name1", email: "eMaIl1"},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				wantErr: false,
			},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: nil},
			createArgs: createArgs{
				user: &domain.User{
					ID:      "",
					Name:    "name1",
					Email:   "email1",
					Version: 0,
				},
			},
			createReturn: createReturn{err: nil},
		},
		{
			name:     "noName",
			args:     args{name: "", email: "email1"},
			expected: expected{user: nil, wantErr: true},
		},
		{
			name:     "noEmail",
			args:     args{name: "name1", email: ""},
			expected: expected{user: nil, wantErr: true},
		},
		{
			name:     "noArgs",
			args:     args{name: "", email: ""},
			expected: expected{user: nil, wantErr: true},
		},
		{
			name:                  "userAlreadyExists",
			args:                  args{name: "name1", email: "email1"},
			expected:              expected{user: nil, wantErr: true},
			findActiveByEmailArgs: findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				err: nil,
			},
		},
		{
			name:                    "findUserFailed",
			args:                    args{name: "name1", email: "email1"},
			expected:                expected{user: nil, wantErr: true},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: errors.New("")},
		},
		{
			name:                    "createFailed",
			args:                    args{name: "name1", email: "email1"},
			expected:                expected{user: nil, wantErr: true},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: nil},
			createArgs: createArgs{
				user: &domain.User{
					ID:      "",
					Name:    "name1",
					Email:   "email1",
					Version: 0,
				},
			},
			createReturn: createReturn{err: errors.New("")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			storageMock := mocks.NewUserStorage(t)
			target := &UsrSvc{storage: storageMock}

			if tt.findActiveByEmailArgs.email != "" {
				storageMock.EXPECT().
					FindActiveByEmail(tt.findActiveByEmailArgs.email).
					Return(tt.findActiveByEmailReturn.user, tt.findActiveByEmailReturn.err)
			}

			if tt.createArgs.user != nil {
				storageMock.EXPECT().
					Create(tt.createArgs.user).
					Run(func(user *domain.User) {
						if !user.Exists() {
							user.ID = "1"
						}
					}).
					Return(tt.createReturn.err)
			}

			user, err := target.Create(tt.args.name, tt.args.email)

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.user, user)
		})
	}
}

func TestService_Disable(t *testing.T) {
	type getArgs struct {
		id string
	}

	type getReturn struct {
		user *domain.User
		err  error
	}

	type updateArgs struct {
		user *domain.User
	}

	type updateReturn struct {
		err error
	}

	type args struct {
		id string
	}

	type expected struct {
		user    *domain.User
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
		getArgs
		getReturn
		updateArgs
		updateReturn
	}{
		{
			name:         "success",
			args:         args{id: "1"},
			expected:     expected{user: &domain.User{ID: "1", Disabled: true, Version: 1}, wantErr: false},
			getArgs:      getArgs{id: "1"},
			getReturn:    getReturn{user: &domain.User{ID: "1", Version: 0}, err: nil},
			updateArgs:   updateArgs{&domain.User{ID: "1", Disabled: true}},
			updateReturn: updateReturn{err: nil},
		},
		{
			name:     "noArg",
			args:     args{id: ""},
			expected: expected{wantErr: true},
		},
		{
			name:      "noUser",
			args:      args{id: "1"},
			expected:  expected{wantErr: true},
			getArgs:   getArgs{id: "1"},
			getReturn: getReturn{user: nil, err: nil},
		},
		{
			name:      "getUserFailed",
			args:      args{id: "1"},
			expected:  expected{wantErr: true},
			getArgs:   getArgs{id: "1"},
			getReturn: getReturn{user: nil, err: errors.New("")},
		},
		{
			name:      "disableFailed",
			args:      args{id: "1"},
			expected:  expected{user: &domain.User{ID: "1", Disabled: true, Version: 0}, wantErr: true},
			getArgs:   getArgs{id: "1"},
			getReturn: getReturn{user: &domain.User{ID: "1", Disabled: true, Version: 0}, err: nil},
		},
		{
			name:         "updateFailed",
			args:         args{id: "1"},
			expected:     expected{user: &domain.User{ID: "1", Disabled: true, Version: 0}, wantErr: true},
			getArgs:      getArgs{id: "1"},
			getReturn:    getReturn{user: &domain.User{ID: "1", Version: 0}, err: nil},
			updateArgs:   updateArgs{&domain.User{ID: "1", Disabled: true}},
			updateReturn: updateReturn{err: errors.New("")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			storageMock := mocks.NewUserStorage(t)
			target := &UsrSvc{storage: storageMock}

			if tt.getArgs.id != "" {
				storageMock.EXPECT().
					Get(tt.getArgs.id).
					Return(tt.getReturn.user, tt.getReturn.err)
			}

			if tt.updateArgs.user != nil {
				storageMock.EXPECT().
					Update(tt.updateArgs.user).
					Run(func(user *domain.User) {
						if user != nil && tt.updateReturn.err == nil {
							user.Version++
						}
					}).
					Return(tt.updateReturn.err)
			}

			err := target.Disable(tt.args.id)
			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.user, tt.getReturn.user)
		})
	}
}

func TestService_Get(t *testing.T) {
	type findActiveByEmailArgs struct {
		email string
	}

	type findActiveByEmailReturn struct {
		user *domain.User
		err  error
	}

	type args struct {
		email string
	}

	type expected struct {
		user    *domain.User
		wantErr bool
	}

	tests := []struct {
		name string
		args
		expected
		findActiveByEmailArgs
		findActiveByEmailReturn
	}{
		{
			name: "success",
			args: args{email: "email1"},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				wantErr: false},
			findActiveByEmailArgs: findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				err: nil},
		},
		{
			name:     "noEmail",
			args:     args{email: ""},
			expected: expected{user: nil, wantErr: true},
		},
		{
			name:                    "findUserFailed",
			args:                    args{email: "email1"},
			expected:                expected{user: nil, wantErr: true},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: errors.New("")},
		},
		{
			name:                    "noUser",
			args:                    args{email: "email1"},
			expected:                expected{user: nil, wantErr: false},
			findActiveByEmailArgs:   findActiveByEmailArgs{email: "email1"},
			findActiveByEmailReturn: findActiveByEmailReturn{user: nil, err: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			storageMock := mocks.NewUserStorage(t)
			target := &UsrSvc{storage: storageMock}

			if tt.findActiveByEmailArgs.email != "" {
				storageMock.EXPECT().
					FindActiveByEmail(tt.findActiveByEmailArgs.email).
					Return(tt.findActiveByEmailReturn.user, tt.findActiveByEmailReturn.err)
			}

			got, err := target.Get(tt.args.email)

			assertions.Equal(tt.expected.user, got)
			assertions.True((err != nil) == tt.expected.wantErr)
		})
	}
}

func TestService_Update(t *testing.T) {
	type getReturn struct {
		user *domain.User
		err  error
	}

	type updateArgs struct {
		user *domain.User
	}

	type updateReturn struct {
		err error
	}

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
		getReturn
		updateArgs
		updateReturn
	}{
		{
			name: "success",
			args: args{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 1,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:       "1",
					Name:     "name1_changed",
					Email:    "email1_changed",
					Disabled: false,
					Version:  2,
				},
				wantErr: false,
			},
			getReturn: getReturn{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				err: nil,
			},
			updateArgs: updateArgs{
				user: &domain.User{
					ID:       "1",
					Name:     "name1_changed",
					Email:    "email1_changed",
					Disabled: false,
					Version:  1,
				},
			},
			updateReturn: updateReturn{err: nil},
		},
		{
			name: "noArg",
			args: args{
				user: nil,
			},
			expected: expected{
				user:    nil,
				wantErr: true,
			},
		},
		{
			name: "noUser",
			args: args{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
				wantErr: true,
			},
			getReturn: getReturn{
				user: nil,
				err:  nil,
			},
		},
		{
			name: "getUserFailed",
			args: args{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
				wantErr: true,
			},
			getReturn: getReturn{
				user: nil,
				err:  errors.New(""),
			},
		},
		{
			name: "disabledUser",
			args: args{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 0,
				},
				wantErr: true,
			},
			getReturn: getReturn{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: true,
					Version:  0,
				},
				err: nil,
			},
		},
		{
			name: "updateFailed",
			args: args{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 1,
				},
			},
			expected: expected{
				user: &domain.User{
					ID:      "1",
					Name:    "name1_changed",
					Email:   "email1_changed",
					Version: 1,
				},
				wantErr: true,
			},
			getReturn: getReturn{
				user: &domain.User{
					ID:       "1",
					Name:     "name1",
					Email:    "email1",
					Disabled: false,
					Version:  0,
				},
				err: nil,
			},
			updateArgs: updateArgs{
				user: &domain.User{
					ID:       "1",
					Name:     "name1_changed",
					Email:    "email1_changed",
					Disabled: false,
					Version:  1,
				},
			},
			updateReturn: updateReturn{err: errors.New("")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)

			storageMock := mocks.NewUserStorage(t)
			target := &UsrSvc{storage: storageMock}

			if tt.args.user.Exists() {
				storageMock.EXPECT().
					Get(tt.args.user.ID).
					Return(tt.getReturn.user, tt.getReturn.err)
			}

			if tt.updateArgs.user.Exists() {
				storageMock.EXPECT().
					Update(tt.updateArgs.user).
					Run(func(user *domain.User) {
						if tt.updateReturn.err == nil {
							user.Version++
						}
					}).
					Return(tt.updateReturn.err)
			}

			err := target.Update(tt.args.user)

			assertions.True((err != nil) == tt.expected.wantErr)
			assertions.Equal(tt.expected.user, tt.args.user)
		})
	}
}
