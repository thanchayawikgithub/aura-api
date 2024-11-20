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

func (suite *ServiceTestSuite) TestAddPost() {

	type args struct {
		ctx context.Context
		req *auraapi.AddPostReq
	}

	testCases := []test.ServiceTestCase[args, *auraapi.AddPostRes]{
		{
			Name: "add post success",
			Mock: func() {
				suite.postStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      2,
					Content: "test",
				}, nil).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddPostReq{
					Content: "test",
				},
			},
			Want: &auraapi.AddPostRes{
				Post: &auradomain.Post{
					ID:      2,
					Content: "test",
				},
			},
			WantErr: false,
			Err:     nil,
		},
		{
			Name: "storage error",
			Mock: func() {
				suite.postStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrInvalidDB).Once()
			},
			Args: args{
				ctx: suite.ctx,
				req: &auraapi.AddPostReq{
					Content: "test",
				},
			},
			Want:    nil,
			WantErr: true,
			Err:     gorm.ErrInvalidDB,
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.Name, func() {
			testCase.Mock()
			got, err := suite.PostService.AddPost(testCase.Args.ctx, testCase.Args.req)
			if testCase.WantErr {
				suite.Error(err)
				suite.Equal(testCase.Err, err)
			} else {
				suite.NoError(err)
				suite.Equal(testCase.Want, got)
			}
		})
	}
}
