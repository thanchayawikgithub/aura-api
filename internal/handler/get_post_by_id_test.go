package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestGetPostByID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.GetPostByIdRes
		wantErr bool
		err     error
	}{
		{
			name: "success",
			mock: func() {
				suite.postStorage.On("WithPreload", mock.Anything, mock.Anything).Return(suite.postStorage).Once()
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			want: &auraapi.GetPostByIdRes{
				Post:     &auradomain.Post{},
				User:     &auradomain.User{},
				Comments: []*auradomain.Comment{},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "post not found",
			mock: func() {
				suite.postStorage.On("WithPreload", mock.Anything, mock.Anything).Return(suite.postStorage).Once()
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrRecordNotFound).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			wantErr: true,
			err:     gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mock()
		})

		got, err := suite.PostService.GetPostByID(tc.args.ctx, tc.args.id)
		if tc.wantErr {
			suite.Error(err)
			suite.Equal(tc.err, err)
		} else {
			suite.NoError(err)
			suite.Equal(tc.want, got)
		}
	}

}
