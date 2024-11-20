// package httpadapter

// import (
// 	"aura/auraapi"
// 	"aura/auradomain"
// 	"aura/internal/pkg/response"
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/mock"
// )

// func (suite *AdapterTestSuite) TestAddUser() {
// 	type args struct {
// 		c echo.Context
// 	}

// 	mockCtx := GenerateMockEchoContext(http.MethodPost, "/v1/user", map[string]interface{}{
// 		"email":        "test@example.com",
// 		"username":     "testuser",
// 		"password":     "password123",
// 		"display_name": "Test User",
// 	})

// 	testCases := []struct {
// 		Name string
// 		Mock func()
// 		Args args
// 		Want error
// 	}{
// 		{
// 			Name: "add user success",
// 			Mock: func() {
// 				suite.userService.On("AddUser", mock.Anything, mock.Anything).Return(&auraapi.AddUserRes{
// 					User: &auradomain.User{
// 						ID:          123,
// 						Email:       "test@example.com",
// 						Username:    "testuser",
// 						DisplayName: "Test User",
// 					},
// 				}, nil).Once()
// 			},
// 			Args: args{
// 				c: mockCtx,
// 			},
// 			Want: response.OK(mockCtx, &auraapi.AddUserRes{
// 				User: &auradomain.User{
// 					ID:          123,
// 					Email:       "test@example.com",
// 					Username:    "testuser",
// 					DisplayName: "Test User",
// 				},
// 			}),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		suite.Run(tc.Name, func() {
// 			if tc.Mock != nil {
// 				tc.Mock()
// 			}

//				err := suite.adapter.AddUser(tc.Args.c)
//				fmt.Println("err", err.Error())
//				suite.ErrorIs(tc.Want, err)
//			})
//		}
//	}
package httpadapter

import (
	"aura/auraapi"
	"aura/auradomain"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
)

func (suite *AdapterTestSuite) TestAddUser() {
	// Create mock context with request body
	mockCtx := GenerateMockEchoContext(http.MethodPost, "/v1/user", map[string]interface{}{
		"email":        "test@example.com",
		"username":     "testuser",
		"password":     "password123",
		"display_name": "Test User",
	})

	// Expected response data
	expectedUser := &auradomain.User{
		ID:          123,
		Email:       "test@example.com",
		Username:    "testuser",
		DisplayName: "Test User",
	}

	// Mock the service response
	suite.userService.On("AddUser", mock.Anything, mock.Anything).Return(&auraapi.AddUserRes{
		User: expectedUser,
	}, nil).Once()

	// Call the handler
	err := suite.adapter.AddUser(mockCtx)

	// Verify response
	suite.NoError(err)

	// Get the response from the recorder
	rec := mockCtx.Response().Writer.(*httptest.ResponseRecorder)

	// Check status code
	suite.Equal(http.StatusCreated, rec.Code)

	// Parse response body
	var response struct {
		Message string           `json:"message"`
		Data    *auradomain.User `json:"data"`
	}
	err = json.NewDecoder(rec.Body).Decode(&response)
	suite.NoError(err)

	// Verify response structure
	suite.Equal("Created", response.Message)
	suite.Equal(expectedUser, response.Data)
}
