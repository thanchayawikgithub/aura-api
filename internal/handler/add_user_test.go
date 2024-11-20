package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	test "aura/tests/app"
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

	testCases := []test.TestCase[args, *auraapi.AddUserRes]{
		{
			Name: "add user success",
			Mock: func() {
				suite.userStorage.On("Save", mock.Anything, mock.Anything).Return(&model.User{
					ID:          3,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddUserReq{
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
					Password:    "password",
				},
			},
			Want: &auraapi.AddUserRes{
				User: &auradomain.User{
					ID:          3,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "storage error",
			Mock: func() {
				suite.userStorage.On("Save", mock.Anything, mock.Anything).Return(&model.User{}, gorm.ErrInvalidDB).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddUserReq{
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
					Password:    "password",
				},
			},
			Want:    &auraapi.AddUserRes{},
			WantErr: true,
			Err:     gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			if tc.Mock != nil {
				tc.Mock()
			}

			got, err := suite.UserService.AddUser(tc.Args.ctx, tc.Args.req)
			if tc.WantErr {
				suite.Error(err)
				suite.Equal(tc.Err, err)
			} else {
				suite.NoError(err)
				suite.Equal(tc.Want, got)
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
