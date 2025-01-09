package requests

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyIDs(t *testing.T) {
	tests := []struct {
		name        string
		propertyIds string
		want        []string
		wantErr     bool
		errorMsg    string
	}{
		{
			name:        "successful case with multiple IDs",
			propertyIds: "123,456,789",
			want:        []string{"123", "456", "789"},
			wantErr:     false,
			errorMsg:    "",
		},
		{
			name:        "successful case with single ID",
			propertyIds: "123",
			want:        []string{"123"},
			wantErr:     false,
			errorMsg:    "",
		},
		{
			name:        "empty property IDs",
			propertyIds: "",
			want:        nil,
			wantErr:     true,
			errorMsg:    "no property IDs provided",
		},
		{
			name:        "handles whitespace",
			propertyIds: "123, 456, 789",
			want:        []string{"123", " 456", " 789"},
			wantErr:     false,
			errorMsg:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create URL with properly encoded query parameters
			reqURL := "/test"
			if tt.propertyIds != "" {
				params := url.Values{}
				params.Add("propertyIds", tt.propertyIds)
				reqURL = "/test?" + params.Encode()
			}

			// Create a new request
			req := httptest.NewRequest("GET", reqURL, nil)
			w := httptest.NewRecorder()

			// Initialize the context properly
			ctx := context.NewContext()
			ctx.Reset(w, req)

			// Create and initialize the controller
			ctrl := &web.Controller{}
			ctrl.Init(ctx, "", "", nil)

			// Call the function being tested
			got, err := GetPropertyIDs(ctrl)

			// Check error cases
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, got)
				return
			}

			// Check success cases
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
