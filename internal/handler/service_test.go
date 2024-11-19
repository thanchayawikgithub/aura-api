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

	storage     *storage.Storage
	userStorage *database.MockUserStorage
	postStorage *database.MockPostStorage

	service     *Service
	UserService *UserService
	PostService *PostService
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
	suite.service = &Service{
		UserStorage: suite.userStorage,
		PostStorage: suite.postStorage,
	}

	suite.UserService = NewUserService(suite.service)
	suite.PostService = NewPostService(suite.service)
}

func (suite *ServiceTestSuite) MockContext() context.Context {
	var tx *gorm.DB

	ctx := context.WithValue(context.TODO(), util.Tx, tx)
	ctx = context.WithValue(ctx, util.UserID, 1)
	ctx = context.WithValue(ctx, util.UserEmail, "test@test.com")

	return ctx
}