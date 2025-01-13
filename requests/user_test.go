package requests

import (
	"beego-api-service/structs"
	"beego-api-service/validation"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Set up test configuration
	err := web.LoadAppConfig("ini", "../conf/app.conf")
	if err != nil {
		panic("Failed to load app.conf for testing: " + err.Error())
	}

	web.BConfig.RunMode = "test"

	// Set any other test-specific configurations
	web.BConfig.WebConfig.AutoRender = false

	// Set test user
	web.AppConfig.Set("TestUser", "srsaurav0")
}

func TestHandleCreateUserRequest_EmptyBody(t *testing.T) {
	controller := &web.Controller{}
	req, err := HandleCreateUserRequest(controller, []byte{})

	assert.Nil(t, req)
	assert.NotNil(t, err)

	// Type assert to *structs.StandardResponse
	errorResponse, ok := err.(*structs.StandardResponse)
	assert.True(t, ok)
	assert.False(t, errorResponse.Success)
	assert.NotNil(t, errorResponse.Error)
	assert.Equal(t, "INVALID_REQUEST", errorResponse.Error.Code)
}

func TestHandleCreateUserRequest(t *testing.T) {
	controller := &web.Controller{}

	tests := []struct {
		name          string
		inputJSON     string
		expectedReq   *validation.CreateUserRequest
		expectedError *structs.StandardResponse
	}{
		{
			name: "Valid request",
			inputJSON: `{
                "name": "Test User",
                "age": 25,
                "email": "test@example.com"
            }`,
			expectedReq: &validation.CreateUserRequest{
				Name:  "Test User",
				Age:   25,
				Email: "test@example.com",
			},
			expectedError: nil,
		},
		{
			name:        "Invalid JSON format",
			inputJSON:   `{"name": "Test User", "email": invalid}`,
			expectedReq: nil,
			expectedError: &structs.StandardResponse{
				Success: false,
				Error: &structs.ErrorInfo{
					Code:    "INVALID_REQUEST",
					Message: "Invalid JSON format: invalid character 'i' looking for beginning of value",
				},
			},
		},
		{
			name: "Validation error - empty name",
			inputJSON: `{
                "name": "",
                "age": 25,
                "email": "test@example.com"
            }`,
			expectedReq: nil,
			expectedError: &structs.StandardResponse{
				Success: false,
				Error: &structs.ErrorInfo{
					Code:    "VALIDATION_ERROR",
					Message: "name is required",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert input JSON to bytes
			body := []byte(tt.inputJSON)

			// Call the function
			req, err := HandleCreateUserRequest(controller, body)

			if tt.expectedError != nil {
				assert.Nil(t, req)
				errorResponse, ok := err.(*structs.StandardResponse)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError.Error.Code, errorResponse.Error.Code)
				assert.Equal(t, tt.expectedError.Success, errorResponse.Success)
			} else {
				assert.NotNil(t, req)
				assert.Equal(t, tt.expectedReq.Name, req.Name)
				assert.Equal(t, tt.expectedReq.Age, req.Age)
				assert.Equal(t, tt.expectedReq.Email, req.Email)
				assert.Nil(t, err)
			}
		})
	}
}
