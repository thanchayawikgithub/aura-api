package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	test "aura/tests/app"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestGetUserByID() {
	type args struct {
		ctx context.Context
		id  uint
	}
	testCases := []test.ServiceTestCase[args, *auraapi.GetUserByIdRes]{
		{
			Name: "success",
			Mock: func() {
				suite.userStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.User{
					ID:          1,
					Email:       "test@test.com",
					Username:    "test",
					Password:    "password",
					DisplayName: "test",
				}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			Want: &auraapi.GetUserByIdRes{
				User: &auradomain.User{
					ID:          1,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "user not found",
			Mock: func() {
				suite.userStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.User{}, gorm.ErrRecordNotFound).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  100,
			},
			Want:    nil,
			WantErr: true,
			Err:     gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			tc.Mock()
			got, err := suite.UserService.GetUserByID(tc.Args.ctx, tc.Args.id)
			if tc.WantErr {
				suite.Error(err)
				suite.Equal(tc.Err, err)
			} else {
				suite.NoError(err)
				suite.Equal(tc.Want, got)
			}
		})
	}
}
