package controllers

import (
	"beego-api-service/models"
	"beego-api-service/validation"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Post(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetByIdentifier(identifier string) (*models.User, error) {
	args := m.Called(identifier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) Update(identifier string, user *models.User) error {
	args := m.Called(identifier, user)
	return args.Error(0)
}

func (m *MockUserService) Delete(identifier string) error {
	args := m.Called(identifier)
	return args.Error(0)
}

func setupTestController() (*UserController, *MockUserService) {
	mockService := new(MockUserService)
	controller := &UserController{
		BaseController: BaseController{},
		userService:    mockService,
	}
	return controller, mockService
}

func initializeContext(req *http.Request, w http.ResponseWriter) *context.Context {
	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input = &context.BeegoInput{
		Context:     ctx,
		RequestBody: []byte{},
	}
	ctx.Output = &context.BeegoOutput{
		Context: ctx,
		Status:  200,
	}
	return ctx
}

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(*MockUserService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "Success",
			requestBody: validation.CreateUserRequest{
				Name:  "John Doe",
				Age:   30,
				Email: "john@example.com",
			},
			setupMock: func(ms *MockUserService) {
				expectedUser := &models.User{
					Name:  "John Doe",
					Age:   30,
					Email: "john@example.com",
				}
				ms.On("Post", mock.MatchedBy(func(user *models.User) bool {
					return user.Name == expectedUser.Name &&
						user.Age == expectedUser.Age &&
						user.Email == expectedUser.Email
				})).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			setupMock:      func(ms *MockUserService) {},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller, mockService := setupTestController()
			w := httptest.NewRecorder()
			tt.setupMock(mockService)

			var reqBody []byte
			var err error
			if tt.requestBody != nil {
				if jsonBody, ok := tt.requestBody.(string); ok {
					reqBody = []byte(jsonBody)
				} else {
					reqBody, err = json.Marshal(tt.requestBody)
					assert.NoError(t, err)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/api/user", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			ctx := initializeContext(req, w)
			controller.Ctx = ctx

			controller.CreateUser()

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name           string
		identifier     string
		setupMock      func(*MockUserService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:       "Success",
			identifier: "1",
			setupMock: func(ms *MockUserService) {
				ms.On("GetByIdentifier", "1").Return(&models.User{
					Name:  "John Doe",
					Age:   30,
					Email: "john@example.com",
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:       "User Not Found",
			identifier: "999",
			setupMock: func(ms *MockUserService) {
				ms.On("GetByIdentifier", "999").Return(nil, errors.New("user not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller, mockService := setupTestController()
			w := httptest.NewRecorder()
			tt.setupMock(mockService)

			req := httptest.NewRequest(http.MethodGet, "/v1/api/user/"+tt.identifier, nil)
			ctx := initializeContext(req, w)
			ctx.Input.SetParam(":identifier", tt.identifier)
			controller.Ctx = ctx

			controller.GetUser()

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {

	tests := []struct {
		name           string
		identifier     string
		requestBody    interface{}
		setupMock      func(*MockUserService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:       "Success",
			identifier: "1",
			requestBody: validation.UpdateUserRequest{
				Name: "John Updated",
				Age:  31,
			},
			setupMock: func(ms *MockUserService) {
				expectedUser := &models.User{
					Name: "John Updated",
					Age:  31,
				}
				ms.On("Update", "1", mock.MatchedBy(func(user *models.User) bool {
					return user.Name == expectedUser.Name &&
						user.Age == expectedUser.Age
				})).Return(nil)
				ms.On("GetByIdentifier", "1").Return(expectedUser, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:       "User Not Found",
			identifier: "999",
			requestBody: validation.UpdateUserRequest{
				Name: "John Updated",
				Age:  31,
			},
			setupMock: func(ms *MockUserService) {
				ms.On("Update", "999", mock.AnythingOfType("*models.User")).Return(errors.New("user not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller, mockService := setupTestController()
			w := httptest.NewRecorder()
			tt.setupMock(mockService)

			jsonBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPut, "/v1/api/user/"+tt.identifier, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			ctx := initializeContext(req, w)
			ctx.Input.SetParam(":identifier", tt.identifier)
			controller.Ctx = ctx

			controller.UpdateUser()

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		identifier     string
		setupMock      func(*MockUserService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:       "Success",
			identifier: "1",
			setupMock: func(ms *MockUserService) {
				ms.On("Delete", "1").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:       "User Not Found",
			identifier: "999",
			setupMock: func(ms *MockUserService) {
				ms.On("Delete", "999").Return(errors.New("user not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller, mockService := setupTestController()
			w := httptest.NewRecorder()
			tt.setupMock(mockService)

			req := httptest.NewRequest(http.MethodDelete, "/v1/api/user/"+tt.identifier, nil)

			ctx := initializeContext(req, w)
			ctx.Input.SetParam(":identifier", tt.identifier)
			controller.Ctx = ctx

			controller.DeleteUser()

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}
