package handler

import (
	"aura/internal/model"
	"aura/internal/util"
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestDeleteComment() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "delete comment success",
			mock: func() {
				userID, err := util.GetUserID(suite.ctx)
				suite.NoError(err)

				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{
					ID:        1,
					UserID:    userID,
					PostID:    1,
					Content:   "test",
					CreatedAt: time.Now(),
				}, nil).Once()
				suite.commentStorage.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "user id not found in context",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			wantErr: true,
			err:     errors.New("user id not found in context"),
		},
		{
			name: "comment not found",
			mock: func() {
				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{}, gorm.ErrRecordNotFound).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			wantErr: true,
			err:     gorm.ErrRecordNotFound,
		},
		{
			name: "no permission",
			mock: func() {
				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{
					ID:        1,
					UserID:    2,
					PostID:    1,
					Content:   "test",
					CreatedAt: time.Now(),
				}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			wantErr: true,
			err:     ErrNoPermission,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.mock != nil {
				tc.mock()
			}

			err := suite.CommentService.DeleteComment(tc.args.ctx, tc.args.id)
			if tc.wantErr {
				suite.Error(err)
				suite.Equal(tc.err, err)
			} else {
				suite.Equal(tc.err, err)
				suite.NoError(err)
			}
		})
	}
}
