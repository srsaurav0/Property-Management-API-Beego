package controllers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/beego/beego/v2/server/web"
// 	"github.com/beego/beego/v2/server/web/context"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetPropertyDetails(t *testing.T) {

// 	// Set mock configuration values
// 	web.AppConfig.Set("externalAPIBaseURL", "http://192.168.0.44:8085/dynamodb-s3-os")

// 	currentTime := "2025-01-09T08:46:06Z" // Updated to provided time
// 	currentUser := "mock"                 // Updated to provided user

// 	tests := []struct {
// 		name           string
// 		propertyID     string
// 		expectedStatus int
// 		wantError      bool
// 		errorMessage   string
// 	}{
// 		{
// 			name:           "Success",
// 			propertyID:     "123456",
// 			expectedStatus: http.StatusOK,
// 			wantError:      false,
// 		},
// 		{
// 			name:           "Empty Property ID",
// 			propertyID:     "",
// 			expectedStatus: http.StatusBadRequest,
// 			wantError:      true,
// 			errorMessage:   "Property ID not provided",
// 		},
// 		{
// 			name:           "Property Not Found",
// 			propertyID:     "invalid",
// 			expectedStatus: http.StatusInternalServerError,
// 			wantError:      true,
// 			errorMessage:   "Failed to fetch property details",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create response recorder
// 			w := httptest.NewRecorder()

// 			// Create new request with mock server URL
// 			r, err := http.NewRequest("GET", "http://localhost:8080/v1/api/property/details/"+tt.propertyID, nil)
// 			assert.NoError(t, err)

// 			// Create a new controller instance
// 			controller := &PropertyDetailsController{}

// 			// Create a new context
// 			ctx := context.NewContext()
// 			ctx.Reset(w, r)
// 			controller.Init(ctx, "", "", nil)

// 			// Set up the property ID in the context
// 			ctx.Input.SetParam(":propertyId", tt.propertyID)

// 			// Set language code
// 			ctx.Input.SetParam("languageCode", "en")

// 			// Execute the handler
// 			controller.GetPropertyDetails()

// 			// Check status code
// 			assert.Equal(t, tt.expectedStatus, w.Code)

// 			// Get response body
// 			responseData := w.Body.Bytes()

// 			// Check response
// 			if tt.wantError {
// 				assert.Contains(t, string(responseData), tt.errorMessage)
// 			} else {
// 				var response map[string]interface{}
// 				err := json.Unmarshal(responseData, &response)
// 				assert.NoError(t, err)

// 				// Verify expected fields
// 				assert.Equal(t, "123456", response["ID"])
// 				assert.Equal(t, float64(1), response["Feed"])
// 				assert.Equal(t, true, response["Published"])

// 				// Check GeoInfo
// 				geoInfo, ok := response["GeoInfo"].(map[string]interface{})
// 				assert.True(t, ok)
// 				assert.Equal(t, "Test City", geoInfo["City"])
// 				assert.Equal(t, "Test Country", geoInfo["Country"])
// 				assert.Equal(t, "TC", geoInfo["CountryCode"])

// 				// Check Property
// 				property, ok := response["Property"].(map[string]interface{})
// 				assert.True(t, ok)
// 				assert.Equal(t, "Test Property", property["PropertyName"])
// 				assert.Equal(t, float64(1000), property["Price"])
// 				assert.Equal(t, currentTime, property["UpdatedAt"])

// 				// Check Counts
// 				counts, ok := property["Counts"].(map[string]interface{})
// 				assert.True(t, ok)
// 				assert.Equal(t, float64(2), counts["Bedroom"])
// 				assert.Equal(t, float64(2), counts["Bathroom"])
// 				assert.Equal(t, float64(10), counts["Reviews"])
// 				assert.Equal(t, float64(4), counts["Occupancy"])

// 				// Check user
// 				assert.Equal(t, currentUser, response["UserLogin"])
// 			}
// 		})
// 	}
// }
