package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"aura/internal/util"
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

	testCases := []struct {
		name    string
		args    args
		mock    func()
		want    *auraapi.EditPostRes
		wantErr bool
		err     error
	}{
		{
			name: "success",
			args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			mock: func() {
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
			wantErr: false,
			err:     nil,
			want: &auraapi.EditPostRes{
				Post: &auradomain.Post{
					ID:      1,
					Content: "change for test",
					UserID:  1,
				},
			},
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
			name: "post not found",
			mock: func() {
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrRecordNotFound).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 20,
			},
			wantErr: true,
			want:    nil,
			err:     gorm.ErrRecordNotFound,
		},
		{
			name: "no permission",
			args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			mock: func() {
				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					UserID:  20,
					Content: "test",
				}, nil).Once()
			},
			wantErr: true,
			want:    nil,
			err:     ErrNoPermission,
		},
		{
			name: "update error",
			args: args{
				ctx: suite.ctx,
				req: &auraapi.EditPostReq{
					Content: "change for test",
				},
				id: 1,
			},
			mock: func() {
				userID, err := util.GetUserID(suite.ctx)
				suite.NoError(err)

				suite.postStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      1,
					UserID:  userID,
					Content: "test",
				}, nil).Once()
				suite.postStorage.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrInvalidDB).Once()
			},
			wantErr: true,
			want:    nil,
			err:     gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.mock != nil {
				tc.mock()
			}

			got, err := suite.PostService.EditPost(tc.args.ctx, tc.args.req, tc.args.id)
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
