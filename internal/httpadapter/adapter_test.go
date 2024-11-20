package httpadapter

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/tests/app/service"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
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

func GenerateMockEchoContext(httpMethod, path string, body map[string]interface{}) echo.Context {
	var reqBody *strings.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		reqBody = strings.NewReader(string(jsonBytes))
	} else {
		reqBody = strings.NewReader("")
	}

	// Create proper request and context
	req := httptest.NewRequest(httpMethod, path, reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = NewCustomValidator()
	ctx := e.NewContext(req, rec)

	return ctx
}
