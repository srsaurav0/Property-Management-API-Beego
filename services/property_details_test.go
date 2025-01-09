package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestFetchPropertyDetails(t *testing.T) {
	// Mock server to simulate external API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request parameters
		assert.Equal(t, "en", r.URL.Query().Get("languageCode"))
		propertyID := r.URL.Query().Get("propertyId")

		switch propertyID {
		case "valid123":
			json.NewEncoder(w).Encode(getMockValidResponse(time.Now().UTC().Format(time.RFC3339)))
		case "invalid456":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid property ID"})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	// Set mock server URL in app config
	web.AppConfig.Set("externalAPIBaseURL", mockServer.URL)

	tests := []struct {
		name           string
		propertyID     string
		expectedError  bool
		validateResult func(*testing.T, structs.PropertyDetailsResponse)
	}{
		{
			name:          "Valid property ID",
			propertyID:    "valid123",
			expectedError: false,
			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
				assert.NotEmpty(t, result.ID)
				assert.NotEmpty(t, result.Property.PropertyName)
			},
		},
		{
			name:          "Invalid property ID",
			propertyID:    "invalid456",
			expectedError: true,
			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
				assert.Empty(t, result)
			},
		},
		{
			name:          "Empty property ID",
			propertyID:    "",
			expectedError: true,
			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
				assert.Empty(t, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FetchPropertyDetails(tt.propertyID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			tt.validateResult(t, result)
		})
	}
}

// Helper function to get mock valid response
func getMockValidResponse(currentTime string) map[string]interface{} {
	return map[string]interface{}{
		"S3": map[string]interface{}{
			"ID":        "TEST123",
			"Feed":      float64(1),
			"Published": true,
			"GeoInfo": map[string]interface{}{
				"Categories": []interface{}{
					map[string]interface{}{
						"Name":       "Category1",
						"Slug":       "cat1",
						"Type":       "type1",
						"Display":    []interface{}{"display1", "display2"},
						"LocationID": "loc123",
					},
				},
				"City":        "Test City",
				"Country":     "Test Country",
				"CountryCode": "TC",
				"Display":     "Test Display",
				"LocationID":  "LOC123",
				"StateAbbr":   "TS",
				"Lat":         "12.345",
				"Lng":         "67.890",
			},
			"Property": map[string]interface{}{
				"Amenities": map[string]interface{}{
					"wifi": "Available",
					"pool": "Available",
				},
				"Counts": map[string]interface{}{
					"Bedroom":   float64(2),
					"Bathroom":  float64(2),
					"Reviews":   float64(10),
					"Occupancy": float64(4),
				},
				"EcoFriendly":  true,
				"FeatureImage": "image.jpg",
				"Image": map[string]interface{}{
					"Count":  float64(2),
					"Images": []interface{}{"img1.jpg", "img2.jpg"},
				},
				"Price":                  float64(100),
				"PropertyName":           "Test Property",
				"PropertySlug":           "test-property",
				"PropertyType":           "Apartment",
				"PropertyTypeCategoryId": "cat123",
				"ReviewScore":            float64(4),
				"ReviewScores": map[string]interface{}{
					"Cleanliness": float64(4.5),
					"Location":    float64(4.2),
				},
				"RoomSize":  float64(50.5),
				"MinStay":   float64(2),
				"UpdatedAt": currentTime,
			},
			"Partner": map[string]interface{}{
				"ID":         "PARTNER123",
				"Archived":   []interface{}{"archived1", "archived2"},
				"OwnerID":    "owner123",
				"HcomID":     "hcom123",
				"BrandId":    "brand123",
				"URL":        "http://example.com",
				"UnitNumber": "unit123",
				"EpCluster":  "cluster1",
			},
		},
	}
}

// func TestTransformData(t *testing.T) {
// 	currentTime := "2025-01-09T06:11:56Z"

// 	tests := []struct {
// 		name           string
// 		input          map[string]interface{}
// 		expectedError  bool
// 		validateResult func(*testing.T, structs.PropertyDetailsResponse)
// 	}{
// 		{
// 			name:          "Valid complete data",
// 			input:         getMockValidResponse(currentTime),
// 			expectedError: false,
// 			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
// 				assert.Equal(t, "TEST123", result.ID)
// 				assert.Equal(t, 1, result.Feed)
// 				assert.True(t, result.Published)

// 				// Validate GeoInfo
// 				assert.Equal(t, "Test City", result.GeoInfo.City)
// 				assert.Equal(t, "Test Country", result.GeoInfo.Country)
// 				assert.Len(t, result.GeoInfo.Categories, 1)

// 				// Validate Property details
// 				assert.Equal(t, "Test Property", result.Property.PropertyName)
// 				assert.Equal(t, 100, result.Property.Price)
// 				assert.Equal(t, 2, result.Property.Counts.Bedroom)
// 				assert.Equal(t, 2, result.Property.Counts.Bathroom)

// 				// Validate Partner details
// 				assert.Equal(t, "PARTNER123", result.Partner.ID)
// 				assert.Contains(t, result.Partner.Archived, "archived1")
// 			},
// 		},
// 		{
// 			name: "Missing S3 data",
// 			input: map[string]interface{}{
// 				"other": "data",
// 			},
// 			expectedError: true,
// 			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
// 				assert.Empty(t, result)
// 			},
// 		},
// 		{
// 			name: "Invalid data types",
// 			input: map[string]interface{}{
// 				"S3": map[string]interface{}{
// 					"ID":        123,    // Invalid: should be string
// 					"Feed":      "1",    // Invalid: should be number
// 					"Published": "true", // Invalid: should be boolean
// 					"GeoInfo": map[string]interface{}{
// 						"City":    456,  // Invalid: should be string
// 						"Country": true, // Invalid: should be string
// 					},
// 					"Property": map[string]interface{}{
// 						"Price": "invalid", // Invalid: should be number
// 					},
// 				},
// 			},
// 			expectedError: true,
// 			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
// 				assert.Empty(t, result)
// 			},
// 		},
// 		{
// 			name: "Nil values in required fields",
// 			input: map[string]interface{}{
// 				"S3": map[string]interface{}{
// 					"ID":        nil,
// 					"Feed":      nil,
// 					"Published": nil,
// 					"GeoInfo":   nil,
// 					"Property":  nil,
// 					"Partner":   nil,
// 				},
// 			},
// 			expectedError: true,
// 			validateResult: func(t *testing.T, result structs.PropertyDetailsResponse) {
// 				assert.Empty(t, result)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var result structs.PropertyDetailsResponse
// 			err := transformData(tt.input, &result)

// 			if tt.expectedError {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}

// 			tt.validateResult(t, result)
// 		})
// 	}
// }
