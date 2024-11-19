package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestAddComment() {
	type args struct {
		ctx context.Context
		req *auraapi.AddCommentReq
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.AddCommentRes
		wantErr bool
		err     error
	}{
		{
			name: "add comment success",
			mock: func() {
				suite.commentStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Comment{
					ID:      1,
					Content: "test",
					PostID:  1,
					UserID:  1,
				}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddCommentReq{
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			want: &auraapi.AddCommentRes{
				Comment: &auradomain.Comment{
					ID:      1,
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "storage error",
			mock: func() {
				suite.commentStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Comment{}, gorm.ErrInvalidData).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddCommentReq{
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			want:    nil,
			wantErr: true,
			err:     gorm.ErrInvalidData,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.mock != nil {
				tc.mock()
			}

			got, err := suite.CommentService.AddComment(tc.args.ctx, tc.args.req)
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
