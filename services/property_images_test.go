package services

import (
	"beego-api-service/structs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestFetchPropertyImages(t *testing.T) {
	// Setup test cases
	tests := []struct {
		name           string
		propertyID     string
		mockResponse   string
		expectedImages structs.ImagesResponse
		expectError    bool
	}{
		{
			name:       "Success with valid images and confidence",
			propertyID: "123",
			mockResponse: `{
                "S3-Gallery": {
                    "category1": [
                        {
                            "label": "bedroom",
                            "url": "http://example.com/image1.jpg",
                            "confidence": 98.5
                        },
                        {
                            "label": "bedroom",
                            "url": "http://example.com/image2.jpg",
                            "confidence": 85.5
                        },
                        {
                            "label": "kitchen",
                            "url": "http://example.com/image3.jpg",
                            "confidence": 97.0
                        }
                    ]
                }
            }`,
			expectedImages: structs.ImagesResponse{
				"bedroom": []string{"http://example.com/image1.jpg"},
				"kitchen": []string{"http://example.com/image3.jpg"},
			},
			expectError: false,
		},
		{
			name:       "Skip images with low confidence",
			propertyID: "124",
			mockResponse: `{
                "S3-Gallery": {
                    "category1": [
                        {
                            "label": "bedroom",
                            "url": "http://example.com/image1.jpg",
                            "confidence": 90.0
                        }
                    ]
                }
            }`,
			expectedImages: structs.ImagesResponse{},
			expectError:    false,
		},
		{
			name:       "Invalid JSON response",
			propertyID: "125",
			mockResponse: `{
                "S3-Gallery": invalid
            }`,
			expectedImages: structs.ImagesResponse{},
			expectError:    true,
		},
		{
			name:       "Missing confidence value",
			propertyID: "126",
			mockResponse: `{
                "S3-Gallery": {
                    "category1": [
                        {
                            "label": "bedroom",
                            "url": "http://example.com/image1.jpg"
                        }
                    ]
                }
            }`,
			expectedImages: structs.ImagesResponse{},
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			// Set mock configuration
			web.AppConfig.Set("externalAPIBaseURL", server.URL)

			// Call the function
			result, err := FetchPropertyImages(tt.propertyID)

			// Assert results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedImages, result)
			}
		})
	}
}

func TestFetchPropertyImagesHTTPError(t *testing.T) {
	// Test case for HTTP request failure
	web.AppConfig.Set("externalAPIBaseURL", "http://invalid-url")

	result, err := FetchPropertyImages("123")

	assert.Error(t, err)
	assert.Empty(t, result)
}
