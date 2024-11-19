package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestGetUserByID() {
	type args struct {
		ctx context.Context
		id  uint
	}
	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.GetUserByIdRes
		wantErr bool
		err     error
	}{
		{
			name: "success",
			mock: func() {
				suite.userStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:       "test@test.com",
					Username:    "test",
					Password:    "password",
					DisplayName: "test",
				}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			want: &auraapi.GetUserByIdRes{
				User: &auradomain.User{
					ID:          1,
					Email:       "test@test.com",
					Username:    "test",
					DisplayName: "test",
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "user not found",
			mock: func() {
				suite.userStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.User{}, gorm.ErrRecordNotFound).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  100,
			},
			want:    nil,
			wantErr: true,
			err:     gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mock()
			got, err := suite.UserService.GetUserByID(tc.args.ctx, tc.args.id)
			if tc.wantErr {
				suite.Error(err)
				suite.Equal(tc.err, err)
			} else {
				suite.NoError(err)
				suite.Equal(tc.want, got)
			}
		})
	}
}
