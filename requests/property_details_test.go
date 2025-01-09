package requests

import (
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func TestGetPropertyID(t *testing.T) {
	tests := []struct {
		name          string
		propertyID    string
		expectedID    string
		expectedError bool
		errorMessage  string
	}{
		{
			name:          "Valid property ID",
			propertyID:    "123456",
			expectedID:    "123456",
			expectedError: false,
			errorMessage:  "",
		},
		{
			name:          "Empty property ID",
			propertyID:    "",
			expectedID:    "",
			expectedError: true,
			errorMessage:  "property ID not provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new controller instance
			ctrl := &web.Controller{}

			// Create a new context
			ctx := &context.Context{
				Input: &context.BeegoInput{},
			}
			ctrl.Init(ctx, "", "", nil)

			// Set up the property ID in the context
			ctx.Input.SetParam(":propertyId", tt.propertyID)

			// Call the function being tested
			gotID, err := GetPropertyID(ctrl)

			// Check error cases
			if tt.expectedError {
				if err == nil {
					t.Errorf("GetPropertyID() expected error but got none")
					return
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("GetPropertyID() error = %v, want error = %v", err.Error(), tt.errorMessage)
					return
				}
			} else {
				if err != nil {
					t.Errorf("GetPropertyID() unexpected error: %v", err)
					return
				}
			}

			// Check the returned ID
			if gotID != tt.expectedID {
				t.Errorf("GetPropertyID() = %v, want %v", gotID, tt.expectedID)
			}
		})
	}
}
