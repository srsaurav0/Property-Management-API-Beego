package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"beego-api-service/structs"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestSendImagesResponse(t *testing.T) {
	tests := []struct {
		name           string
		inputData      structs.ImagesResponse
		expectedStatus int
		wantError      bool
	}{
		{
			name: "Success - Valid Images Response",
			inputData: structs.ImagesResponse{
				"primary": []string{
					"https://example.com/image1.jpg",
					"https://example.com/image2.jpg",
				},
				"secondary": []string{
					"https://example.com/image3.jpg",
					"https://example.com/image4.jpg",
				},
			},
			expectedStatus: http.StatusOK,
			wantError:      false,
		},
		{
			name:           "Success - Empty Images Response",
			inputData:      structs.ImagesResponse{},
			expectedStatus: http.StatusOK,
			wantError:      false,
		},
		{
			name: "Success - Single Category Images",
			inputData: structs.ImagesResponse{
				"thumbnail": []string{
					"https://example.com/thumb1.jpg",
				},
			},
			expectedStatus: http.StatusOK,
			wantError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create response recorder
			w := httptest.NewRecorder()

			// Create new request
			r := httptest.NewRequest("GET", "/api/images", nil)

			// Create a new controller instance
			controller := &web.Controller{}

			// Create and setup context
			ctx := context.NewContext()
			ctx.Reset(w, r)
			controller.Init(ctx, "", "", nil)

			// Call the function
			SendImagesResponse(controller, tt.inputData)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.wantError {
				// Verify response body
				var response structs.ImagesResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Verify the response matches input data
				assert.Equal(t, len(tt.inputData), len(response))

				// Check each category and its images
				for category, expectedUrls := range tt.inputData {
					actualUrls, exists := response[category]
					assert.True(t, exists, "Category %s should exist in response", category)
					assert.Equal(t, expectedUrls, actualUrls, "Images for category %s should match", category)
				}
			}
		})
	}
}
