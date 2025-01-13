package validation

import (
	"reflect"
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

func TestCreateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name          string
		request       CreateUserRequest
		expectedError string
	}{
		{
			name: "Valid request",
			request: CreateUserRequest{
				Name:  "Test User",
				Age:   25,
				Email: "test@example.com",
			},
			expectedError: "",
		},
		{
			name: "Empty name",
			request: CreateUserRequest{
				Name:  "",
				Age:   25,
				Email: "test@example.com",
			},
			expectedError: "name is required",
		},
		{
			name: "Zero age",
			request: CreateUserRequest{
				Name:  "Test User",
				Age:   0,
				Email: "test@example.com",
			},
			expectedError: "invalid age",
		},
		{
			name: "Negative age",
			request: CreateUserRequest{
				Name:  "Test User",
				Age:   -1,
				Email: "test@example.com",
			},
			expectedError: "invalid age",
		},
		{
			name: "Empty email",
			request: CreateUserRequest{
				Name:  "Test User",
				Age:   25,
				Email: "",
			},
			expectedError: "email is required",
		},
		{
			name: "All fields empty",
			request: CreateUserRequest{
				Name:  "",
				Age:   0,
				Email: "",
			},
			expectedError: "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}

func TestUpdateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name          string
		request       UpdateUserRequest
		expectedError string
	}{
		{
			name: "Valid request with both fields",
			request: UpdateUserRequest{
				Name: "Test User",
				Age:  25,
			},
			expectedError: "",
		},
		{
			name: "Valid request with only name",
			request: UpdateUserRequest{
				Name: "Test User",
			},
			expectedError: "",
		},
		{
			name: "Valid request with only age",
			request: UpdateUserRequest{
				Age: 25,
			},
			expectedError: "",
		},
		{
			name: "Empty request",
			request: UpdateUserRequest{
				Name: "",
				Age:  0,
			},
			expectedError: "at least one field (name or age) must be provided for update",
		},
		{
			name: "Negative age",
			request: UpdateUserRequest{
				Name: "Test User",
				Age:  -1,
			},
			expectedError: "invalid age",
		},
		{
			name: "Zero age with name",
			request: UpdateUserRequest{
				Name: "Test User",
				Age:  0,
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}

// TestCreateUserRequest_Fields tests the struct field tags
func TestCreateUserRequest_Fields(t *testing.T) {
	request := CreateUserRequest{}

	t.Run("Name field tags", func(t *testing.T) {
		field, ok := reflect.TypeOf(request).FieldByName("Name")
		assert.True(t, ok)
		assert.Equal(t, "name", field.Tag.Get("json"))
		assert.Equal(t, "Required", field.Tag.Get("valid"))
	})

	t.Run("Age field tags", func(t *testing.T) {
		field, ok := reflect.TypeOf(request).FieldByName("Age")
		assert.True(t, ok)
		assert.Equal(t, "age", field.Tag.Get("json"))
		assert.Equal(t, "Required", field.Tag.Get("valid"))
	})

	t.Run("Email field tags", func(t *testing.T) {
		field, ok := reflect.TypeOf(request).FieldByName("Email")
		assert.True(t, ok)
		assert.Equal(t, "email", field.Tag.Get("json"))
		assert.Equal(t, "Required;Email", field.Tag.Get("valid"))
	})
}

// TestUpdateUserRequest_Fields tests the struct field tags
func TestUpdateUserRequest_Fields(t *testing.T) {
	request := UpdateUserRequest{}

	t.Run("Name field tags", func(t *testing.T) {
		field, ok := reflect.TypeOf(request).FieldByName("Name")
		assert.True(t, ok)
		assert.Equal(t, "name,omitempty", field.Tag.Get("json"))
	})

	t.Run("Age field tags", func(t *testing.T) {
		field, ok := reflect.TypeOf(request).FieldByName("Age")
		assert.True(t, ok)
		assert.Equal(t, "age,omitempty", field.Tag.Get("json"))
	})
}
