package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestGetCommentByID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.GetCommentByIdRes
		wantErr bool
		err     error
	}{
		{
			name: "success",
			mock: func() {
				suite.commentStorage.On("FindByID", mock.Anything, mock.Anything).Return(&model.Comment{}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			want: &auraapi.GetCommentByIdRes{
				Comment: &auradomain.Comment{},
			},
			wantErr: false,
			err:     nil,
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
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mock()
		})

		got, err := suite.CommentService.GetCommentByID(tc.args.ctx, tc.args.id)
		if tc.wantErr {
			suite.Error(err)
			suite.Equal(tc.err, err)
		} else {
			suite.NoError(err)
			suite.Equal(tc.want, got)
		}
	}
}
