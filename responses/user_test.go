package responses

import (
	"beego-api-service/structs"
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

func TestNewSuccessResponse(t *testing.T) {
	tests := []struct {
		name           string
		inputData      interface{}
		inputMessage   string
		expectedResult *structs.StandardResponse
	}{
		{
			name:         "Success response with string data",
			inputData:    "test data",
			inputMessage: "Operation successful",
			expectedResult: &structs.StandardResponse{
				Success: true,
				Data:    "test data",
				Message: "Operation successful",
			},
		},
		{
			name:         "Success response with struct data",
			inputData:    struct{ ID int }{ID: 1},
			inputMessage: "Data retrieved",
			expectedResult: &structs.StandardResponse{
				Success: true,
				Data:    struct{ ID int }{ID: 1},
				Message: "Data retrieved",
			},
		},
		{
			name:         "Success response with nil data",
			inputData:    nil,
			inputMessage: "Empty response",
			expectedResult: &structs.StandardResponse{
				Success: true,
				Data:    nil,
				Message: "Empty response",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewSuccessResponse(tt.inputData, tt.inputMessage)

			assert.NotNil(t, result)
			assert.True(t, result.Success)
			assert.Equal(t, tt.expectedResult.Data, result.Data)
			assert.Equal(t, tt.expectedResult.Message, result.Message)
			assert.Nil(t, result.Error)
		})
	}
}

func TestNewErrorResponse(t *testing.T) {
	tests := []struct {
		name           string
		inputCode      string
		inputMessage   string
		expectedResult *structs.StandardResponse
	}{
		{
			name:         "Error response with validation error",
			inputCode:    "VALIDATION_ERROR",
			inputMessage: "Invalid input",
			expectedResult: &structs.StandardResponse{
				Success: false,
				Error: &structs.ErrorInfo{
					Code:    "VALIDATION_ERROR",
					Message: "Invalid input",
				},
			},
		},
		{
			name:         "Error response with not found error",
			inputCode:    "NOT_FOUND",
			inputMessage: "Resource not found",
			expectedResult: &structs.StandardResponse{
				Success: false,
				Error: &structs.ErrorInfo{
					Code:    "NOT_FOUND",
					Message: "Resource not found",
				},
			},
		},
		{
			name:         "Error response with server error",
			inputCode:    "SERVER_ERROR",
			inputMessage: "Internal server error",
			expectedResult: &structs.StandardResponse{
				Success: false,
				Error: &structs.ErrorInfo{
					Code:    "SERVER_ERROR",
					Message: "Internal server error",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewErrorResponse(tt.inputCode, tt.inputMessage)

			assert.NotNil(t, result)
			assert.False(t, result.Success)
			assert.NotNil(t, result.Error)
			assert.Equal(t, tt.expectedResult.Error.Code, result.Error.Code)
			assert.Equal(t, tt.expectedResult.Error.Message, result.Error.Message)
			assert.Nil(t, result.Data)
			assert.Empty(t, result.Message)
		})
	}
}

// TestResponseStructure tests that the response structures are properly initialized
func TestResponseStructure(t *testing.T) {
	t.Run("Success response structure", func(t *testing.T) {
		response := NewSuccessResponse("test", "test message")
		assert.NotNil(t, response)
		assert.True(t, response.Success)
		assert.NotNil(t, response.Data)
		assert.NotEmpty(t, response.Message)
		assert.Nil(t, response.Error)
	})

	t.Run("Error response structure", func(t *testing.T) {
		response := NewErrorResponse("ERROR_CODE", "error message")
		assert.NotNil(t, response)
		assert.False(t, response.Success)
		assert.Nil(t, response.Data)
		assert.Empty(t, response.Message)
		assert.NotNil(t, response.Error)
		assert.NotEmpty(t, response.Error.Code)
		assert.NotEmpty(t, response.Error.Message)
	})
}
