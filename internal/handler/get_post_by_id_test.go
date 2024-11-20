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

func (suite *ServiceTestSuite) TestGetPostByID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []test.ServiceTestCase[args, *auraapi.GetPostByIdRes]{
		{
			Name: "success",
			Mock: func() {
				suite.postStorage.On("WithPreload", mock.Anything, mock.Anything).Return(suite.postStorage).Once()
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			Want: &auraapi.GetPostByIdRes{
				Post:     &auradomain.Post{},
				User:     &auradomain.User{},
				Comments: []*auradomain.Comment{},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "post not found",
			Mock: func() {
				suite.postStorage.On("WithPreload", mock.Anything, mock.Anything).Return(suite.postStorage).Once()
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrRecordNotFound).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			WantErr: true,
			Err:     gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			if tc.Mock != nil {
				tc.Mock()
			}

			got, err := suite.PostService.GetPostByID(tc.Args.ctx, tc.Args.id)
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
