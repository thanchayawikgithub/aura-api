package handler

import (
	"aura/auraapi"
	"aura/auradomain"
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func (suite *ServiceTestSuite) TestAddPost() {

	type args struct {
		ctx context.Context
		req *auraapi.AddPostReq
	}

	testCases := []struct {
		name    string
		mock    func()
		args    args
		want    *auraapi.AddPostRes
		wantErr bool
		err     error
	}{
		{
			name: "add post success",
			mock: func() {
				suite.postStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Post{
					ID:      2,
					Content: "test",
					UserID:  1,
				}, nil).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddPostReq{
					Content: "test",
				},
			},
			want: &auraapi.AddPostRes{
				Post: &auradomain.Post{
					ID:      2,
					Content: "test",
					UserID:  1,
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "storage error",
			mock: func() {
				suite.postStorage.On("Save", mock.Anything, mock.Anything).Return(&model.Post{}, gorm.ErrInvalidDB).Once()
			},
			args: args{
				ctx: suite.ctx,
				req: &auraapi.AddPostReq{
					Content: "test",
				},
			},
			want:    nil,
			wantErr: true,
			err:     gorm.ErrInvalidDB,
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			testCase.mock()
			got, err := suite.PostService.AddPost(testCase.args.ctx, testCase.args.req)
			if testCase.wantErr {
				suite.Error(err)
				suite.Equal(testCase.err, err)
			} else {
				suite.NoError(err)
				suite.Equal(testCase.want, got)
			}
		})
	}
}
