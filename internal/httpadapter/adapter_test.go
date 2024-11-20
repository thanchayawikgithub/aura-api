package httpadapter

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/tests/app/service"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdapterTestSuite struct {
	suite.Suite
	cfg         *config.Config
	adapter     *Adapter
	userService *service.MockUserService
}

func TestAdapterTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AdapterTestSuite))
}

func (suite *AdapterTestSuite) SetupSuite() {
	suite.cfg = new(config.Config)
	suite.userService = new(service.MockUserService)
	suite.adapter = &Adapter{
		cfg:         suite.cfg,
		service:     new(handler.Service),
		userService: suite.userService,
	}
}
