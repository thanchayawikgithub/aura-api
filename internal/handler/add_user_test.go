package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestAddUser() {
	type args struct {
		ctx context.Context
		req *auraapi.AddUserReq
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.AddUserRes
		wantErr bool
		err     error
	}{
		{
			name: "add user success",
			mock: func() {
				suite.userStorage.On("Save", mock.Anything, mock.Anything).Return(&model.User{
					ID:          3,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddUserReq{
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
					Password:    "password",
				},
			},
			want: &auraapi.AddUserRes{
				User: &auradomain.User{
					ID:          3,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "storage error",
			mock: func() {
				suite.userStorage.On("Save", mock.Anything, mock.Anything).Return(&model.User{}, gorm.ErrInvalidDB).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddUserReq{
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
					Password:    "password",
				},
			},
			want:    &auraapi.AddUserRes{},
			wantErr: true,
			err:     gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.mock != nil {
				tc.mock()
			}

			got, err := suite.UserService.AddUser(tc.args.ctx, tc.args.req)
			if tc.wantErr {
				suite.Error(err)
				suite.Equal(tc.err, err)
			} else {
				suite.NoError(err)
				suite.Equal(tc.want, got)
			}
		})
	}

	suite.Run("password too long", func() {
		got, err := suite.UserService.AddUser(suite.ctx, &auraapi.AddUserReq{
			Email:       "test@test.com",
			Username:    "test",
			DisplayName: "test",
			Password:    string(make([]byte, 73)),
		})

		suite.Nil(got)
		suite.Error(err)
		suite.Equal(bcrypt.ErrPasswordTooLong, err)
	})
}
