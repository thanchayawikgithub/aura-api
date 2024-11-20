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

func (suite *ServiceTestSuite) TestGetPostsByUserID() {
	type args struct {
		ctx context.Context
		id  uint
	}

	testCases := []test.TestCase[args, *auraapi.GetPostsByUserIDRes]{
		{
			Name: "success",
			Mock: func() {
				suite.postStorage.On("FindByUserID", mock.Anything, mock.Anything).Return([]*model.Post{}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
			},
			Want: &auraapi.GetPostsByUserIDRes{
				Posts: []*auradomain.Post{},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "error",
			Mock: func() {
				suite.postStorage.On("FindByUserID", mock.Anything, mock.Anything).Return([]*model.Post{}, gorm.ErrInvalidDB).Once()
			},
			Args: args{
				ctx: suite.ctx,
				id:  1,
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

			got, err := suite.PostService.GetPostsByUserID(tc.Args.ctx, tc.Args.id)
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
