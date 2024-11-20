package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"aura/internal/util"
	test "aura/tests/app"
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestEditPost() {
	type args struct {
		ctx context.Context
		req *auraapi.EditPostReq
		id  uint
	}

	testCases := []test.ServiceTestCase[args, *auraapi.EditPostRes]{
		{
			Name: "success",
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			Mock: func() {
				userID, err := util.GetUserID(suite.ctx)
				suite.NoError(err)

				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					Content: "test",
					UserID:  userID,
				}, nil).Once()
				suite.postStorage.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					Content: "change for test",
					UserID:  userID,
				}, nil).Once()
			},
			WantErr: false,
			Err:     nil,
			Want: &auraapi.EditPostRes{
				Post: &auradomain.Post{
					ID:      1,
					Content: "change for test",
					UserID:  1,
				},
			},
		},
		{
			Name: "user id not found in context",
			Args: args{
				ctx: context.TODO(),
				id:  1,
			},
			WantErr: true,
			Err:     errors.New("user id not found in context"),
		},
		{
			Name: "post not found",
			Mock: func() {
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrRecordNotFound).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 20,
			},
			WantErr: true,
			Err:     gorm.ErrRecordNotFound,
		},
		{
			Name: "no permission",
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			Mock: func() {
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					UserID:  20,
					Content: "test",
				}, nil).Once()
			},
			WantErr: true,
			Err:     ErrNoPermission,
		},
		{
			Name: "update error",
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			Mock: func() {
				userID, err := util.GetUserID(suite.ctx)
				suite.NoError(err)

				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					UserID:  userID,
					Content: "test",
				}, nil).Once()
				suite.postStorage.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrInvalidDB).Once()
			},
			WantErr: true,
			Err:     gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			if tc.Mock != nil {
				tc.Mock()
			}

			got, err := suite.PostService.EditPost(tc.Args.ctx, tc.Args.req, tc.Args.id)
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
