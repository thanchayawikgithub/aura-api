package handler

import (
	"aura/internal/config"
	"aura/internal/storage"
	"aura/internal/util"
	"aura/tests/app/database"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceTestSuite struct {
	suite.Suite
	cfg *config.Config
	ctx context.Context

	storage        *storage.Storage
	userStorage    *database.MockUserStorage
	postStorage    *database.MockPostStorage
	commentStorage *database.MockCommentStorage

	service        *Service
	UserService    *UserService
	PostService    *PostService
	CommentService *CommentService
}

func TestServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupSuite() {
	suite.cfg = &config.Config{}
	suite.ctx = suite.MockContext()

	suite.storage = &storage.Storage{}
	suite.userStorage = new(database.MockUserStorage)
	suite.postStorage = new(database.MockPostStorage)
	suite.commentStorage = new(database.MockCommentStorage)

	suite.service = &Service{
		UserStorage:    suite.userStorage,
		PostStorage:    suite.postStorage,
		CommentStorage: suite.commentStorage,
	}

	suite.UserService = NewUserService(suite.service)
	suite.PostService = NewPostService(suite.service)
	suite.CommentService = NewCommentService(suite.service)
}

func (suite *ServiceTestSuite) MockContext() context.Context {
	var tx *gorm.DB

	ctx := context.WithValue(context.TODO(), util.Tx, tx)
	ctx = context.WithValue(ctx, util.UserID, uint(1))
	ctx = context.WithValue(ctx, util.UserEmail, "test@test.com")

	return ctx
}

// func (suite *ServiceTestSuite) TearDownSuite() {
// 	// Reset all mock storages
// 	suite.userStorage = nil
// 	suite.postStorage = nil
// 	suite.commentStorage = nil
// 	suite.storage = nil

// 	// Clear services
// 	suite.service = nil
// 	suite.UserService = nil
// 	suite.PostService = nil
// 	suite.CommentService = nil

// 	// Clear context and config
// 	suite.ctx = nil
// 	suite.cfg = nil
// }
