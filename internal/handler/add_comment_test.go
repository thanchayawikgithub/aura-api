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

func (suite *ServiceTestSuite) TestAddComment() {
	type args struct {
		ctx context.Context
		req *auraapi.AddCommentReq
	}

	testCases := []test.TestCase[args, *auraapi.AddCommentRes]{
		{
			Name: "add comment success",
			Mock: func() {
				suite.commentStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Comment{
					ID:      1,
					Content: "test",
					PostID:  1,
					UserID:  1,
				}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddCommentReq{
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			Want: &auraapi.AddCommentRes{
				Comment: &auradomain.Comment{
					ID:      1,
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "storage error",
			Mock: func() {
				suite.commentStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Comment{}, gorm.ErrInvalidData).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddCommentReq{
					Content: "test",
					PostID:  1,
					UserID:  1,
				},
			},
			Want:    nil,
			WantErr: true,
			Err:     gorm.ErrInvalidData,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			if tc.Mock != nil {
				tc.Mock()
			}

			got, err := suite.CommentService.AddComment(tc.Args.ctx, tc.Args.req)
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
