package handler

import (
	"aura/internal/model"
	"aura/internal/util"
	test "aura/tests/app"
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestDeletePost() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []test.TestCase[args, error]{
		{
			Name: "success",
			Mock: func() {
				userID, err := util.GetUserID(suite.ctx)
				suite.NoError(err)

				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					UserID: userID,
				}, nil).Once()
				suite.postStorage.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			WantErr: false,
			Err:     nil,
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
				id:  1,
			},
			WantErr: true,
			Err:     gorm.ErrRecordNotFound,
		},
		{
			Name: "no permission",
			Mock: func() {
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:     1,
					UserID: 2,
				}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			WantErr: true,
			Err:     ErrNoPermission,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			if tc.Mock != nil {
				tc.Mock()
			}

			err := suite.PostService.DeletePost(tc.Args.ctx, tc.Args.id)
			if tc.WantErr {
				suite.Error(err)
				suite.Equal(tc.Err, err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
