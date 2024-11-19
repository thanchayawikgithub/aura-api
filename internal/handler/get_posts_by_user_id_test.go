package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestGetPostsByUserID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.GetPostsByUserIDRes
		wantErr bool
		err     error
	}{
		{
			name: "success",
			mock: func() {
				suite.postStorage.On("FindByUserID", mock.Anything, mock.Anything).Return([]*model.Post{}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			want: &auraapi.GetPostsByUserIDRes{
				Posts: []*auradomain.Post{},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "error",
			mock: func() {
				suite.postStorage.On("FindByUserID", mock.Anything, mock.Anything).Return([]*model.Post{}, gorm.ErrInvalidDB).Once()
			},
			args: args{
				ctx: suite.ctx,
				id:  1,
			},
			wantErr: true,
			err:     gorm.ErrInvalidDB,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.mock != nil {
				tc.mock()
			}

			got, err := suite.PostService.GetPostsByUserID(tc.args.ctx, tc.args.id)
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
