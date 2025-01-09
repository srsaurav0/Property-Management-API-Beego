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

func TestSendPropertyDetailsResponses(t *testing.T) {
	tests := []struct {
		name           string
		inputData      []structs.PropertyDetailsResponse
		expectedStatus int
	}{
		{
			name: "successful response with multiple properties",
			inputData: []structs.PropertyDetailsResponse{
				createMockPropertyResponse("123", true),
				createMockPropertyResponse("456", false),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "successful response with empty array",
			inputData:      []structs.PropertyDetailsResponse{},
			expectedStatus: http.StatusOK,
		},
		{
			name: "successful response with single property",
			inputData: []structs.PropertyDetailsResponse{
				createMockPropertyResponse("789", true),
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request and response recorder
			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			// Initialize the context
			ctx := context.NewContext()
			ctx.Reset(w, req)

			// Create and initialize the controller
			ctrl := &web.Controller{}
			ctrl.Init(ctx, "", "", nil)

			// Call the function being tested
			SendPropertyDetailsResponses(ctrl, tt.inputData)

			// Get the response
			resp := w.Result()

			// Check the status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// For successful responses, verify the JSON body
			if tt.expectedStatus == http.StatusOK {
				// Decode the response body
				var decodedResponse []structs.PropertyDetailsResponse
				err := json.NewDecoder(w.Body).Decode(&decodedResponse)
				assert.NoError(t, err)

				// Compare the decoded response with input data
				assert.Equal(t, len(tt.inputData), len(decodedResponse))

				if len(tt.inputData) > 0 {
					// Verify specific fields from the first item
					assert.Equal(t, tt.inputData[0].ID, decodedResponse[0].ID)
					assert.Equal(t, tt.inputData[0].Published, decodedResponse[0].Published)
					assert.Equal(t, tt.inputData[0].Property.PropertyName, decodedResponse[0].Property.PropertyName)
				}

				// Verify Content-Type header
				assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
			}
		})
	}
}

// createMockPropertyResponse creates a mock PropertyDetailsResponse with realistic test data
func createMockPropertyResponse(id string, published bool) structs.PropertyDetailsResponse {
	response := structs.PropertyDetailsResponse{
		ID:        id,
		Feed:      1,
		Published: published,
		GeoInfo: struct {
			Categories []struct {
				Name       string   `json:"Name"`
				Slug       string   `json:"Slug"`
				Type       string   `json:"Type"`
				Display    []string `json:"Display"`
				LocationID string   `json:"LocationID"`
			} `json:"Categories"`
			City        string `json:"City"`
			Country     string `json:"Country"`
			CountryCode string `json:"CountryCode"`
			Display     string `json:"Display"`
			LocationID  string `json:"LocationID"`
			StateAbbr   string `json:"StateAbbr"`
			Lat         string `json:"Lat"`
			Lng         string `json:"Lng"`
		}{
			Categories: []struct {
				Name       string   `json:"Name"`
				Slug       string   `json:"Slug"`
				Type       string   `json:"Type"`
				Display    []string `json:"Display"`
				LocationID string   `json:"LocationID"`
			}{
				{
					Name:       "Test Category",
					Slug:       "test-category",
					Type:       "location",
					Display:    []string{"Display 1", "Display 2"},
					LocationID: "loc_" + id,
				},
			},
			City:        "Test City",
			Country:     "Test Country",
			CountryCode: "TC",
			Display:     "Test Display",
			LocationID:  "loc_" + id,
			StateAbbr:   "TS",
			Lat:         "12.345",
			Lng:         "67.890",
		},
	}

	// Initialize Property struct
	response.Property.Amenities = map[string]string{
		"wifi": "Yes",
		"pool": "No",
	}
	response.Property.Counts.Bedroom = 2
	response.Property.Counts.Bathroom = 1
	response.Property.Counts.Reviews = 10
	response.Property.Counts.Occupancy = 4
	response.Property.EcoFriendly = true
	response.Property.FeatureImage = "feature_" + id + ".jpg"
	response.Property.Image = &struct {
		Count  int      `json:"Count,omitempty"`
		Images []string `json:"Images,omitempty"`
	}{
		Count:  2,
		Images: []string{"image1.jpg", "image2.jpg"},
	}
	response.Property.Price = 150
	response.Property.PropertyName = "Test Property " + id
	response.Property.PropertySlug = "test-property-" + id
	response.Property.PropertyType = "Apartment"
	response.Property.PropertyTypeCategoryId = "apt_" + id
	response.Property.ReviewScore = 85
	response.Property.ReviewScores = map[string]float64{
		"cleanliness": 4.5,
		"location":    4.8,
	}
	response.Property.RoomSize = 45.5
	response.Property.MinStay = 2
	response.Property.UpdatedAt = "2024-01-09T12:00:00Z"

	// Initialize Partner struct
	response.Partner.ID = "partner_" + id
	response.Partner.Archived = []string{"2023-12-01"}
	response.Partner.OwnerID = "owner_" + id
	response.Partner.HcomID = "hcom_" + id
	response.Partner.BrandId = "brand_" + id
	response.Partner.URL = "https://example.com/property/" + id
	response.Partner.UnitNumber = "unit_" + id
	response.Partner.EpCluster = "cluster_1"

	return response
}
