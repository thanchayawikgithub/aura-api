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

func (suite *ServiceTestSuite) TestGetCommentByID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []test.ServiceTestCase[args, *auraapi.GetCommentByIdRes]{
		{
			Name: "Success",
			Mock: func() {
				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			Want: &auraapi.GetCommentByIdRes{
				Comment: &auradomain.Comment{},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "comment not found",
			Mock: func() {
				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{}, gorm.ErrRecordNotFound).Once()
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

			got, err := suite.CommentService.GetCommentByID(tc.Args.ctx, tc.Args.id)
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
