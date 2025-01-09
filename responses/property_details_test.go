package responses

import (
	"beego-api-service/structs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestSendPropertyDetailsResponse(t *testing.T) {
	tests := []struct {
		name           string
		input          structs.PropertyDetailsResponse
		expectedStatus int
	}{
		{
			name:           "Complete property details",
			input:          getCompletePropertyDetails(),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Property details without optional fields",
			input:          getMinimalPropertyDetails(),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Empty property details",
			input:          structs.PropertyDetailsResponse{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test context
			w := httptest.NewRecorder()
			context := context.NewContext()
			context.Reset(w, httptest.NewRequest("GET", "/test", nil))

			// Create a test controller
			controller := web.Controller{}
			controller.Init(context, "", "", nil)

			// Call the function
			SendPropertyDetailsResponse(&controller, tt.input)

			// Assert response status
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Verify the JSON content
			var response structs.PropertyDetailsResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.input, response)

			// Additional validation for specific fields
			if tt.name == "Complete property details" {
				assert.NotNil(t, response.Property.Image)
				assert.NotEmpty(t, response.Property.ReviewScores)
				assert.NotEmpty(t, response.Property.Amenities)
				assert.NotEmpty(t, response.GeoInfo.Categories)
			}
		})
	}
}

func TestSendErrorResponse(t *testing.T) {
	tests := []struct {
		name           string
		message        string
		status         int
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Not Found Error",
			message:        "Property not found",
			status:         http.StatusNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Property not found",
		},
		{
			name:           "Bad Request Error",
			message:        "Invalid property ID",
			status:         http.StatusBadRequest,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid property ID",
		},
		{
			name:           "Internal Server Error",
			message:        "Failed to fetch property details",
			status:         http.StatusInternalServerError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to fetch property details",
		},
		{
			name:           "Empty Error Message",
			message:        "",
			status:         http.StatusBadRequest,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test context
			w := httptest.NewRecorder()
			context := context.NewContext()
			context.Reset(w, httptest.NewRequest("GET", "/test", nil))

			// Create a test controller
			controller := web.Controller{}
			controller.Init(context, "", "", nil)

			// Call the function
			SendErrorResponse(&controller, tt.message, tt.status)

			// Assert response status and body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

// Helper function to create a complete property details response
func getCompletePropertyDetails() structs.PropertyDetailsResponse {
	return structs.PropertyDetailsResponse{
		ID:        "PROP123",
		Feed:      1,
		Published: true,
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
					Name:       "Luxury",
					Slug:       "luxury-homes",
					Type:       "category",
					Display:    []string{"Luxury Homes", "Premium Properties"},
					LocationID: "LOC001",
				},
			},
			City:        "New York",
			Country:     "United States",
			CountryCode: "US",
			Display:     "New York, United States",
			LocationID:  "NYC123",
			StateAbbr:   "NY",
			Lat:         "40.7128",
			Lng:         "-74.0060",
		},
		Property: struct {
			Amenities map[string]string `json:"Amenities"`
			Counts    struct {
				Bedroom   int `json:"Bedroom"`
				Bathroom  int `json:"Bathroom"`
				Reviews   int `json:"Reviews"`
				Occupancy int `json:"Occupancy"`
			} `json:"Counts"`
			EcoFriendly  bool   `json:"EcoFriendly"`
			FeatureImage string `json:"FeatureImage"`
			Image        *struct {
				Count  int      `json:"Count,omitempty"`
				Images []string `json:"Images,omitempty"`
			} `json:"Image,omitempty"`
			Price                  int                `json:"Price"`
			PropertyName           string             `json:"PropertyName"`
			PropertySlug           string             `json:"PropertySlug"`
			PropertyType           string             `json:"PropertyType"`
			PropertyTypeCategoryId string             `json:"PropertyTypeCategoryId"`
			ReviewScore            int                `json:"ReviewScore"`
			ReviewScores           map[string]float64 `json:"ReviewScores,omitempty"`
			RoomSize               float64            `json:"RoomSize"`
			MinStay                int                `json:"MinStay"`
			UpdatedAt              string             `json:"UpdatedAt"`
		}{
			Amenities: map[string]string{
				"wifi":    "Free WiFi",
				"pool":    "Swimming Pool",
				"parking": "Free Parking",
				"aircon":  "Air Conditioning",
			},
			Counts: struct {
				Bedroom   int `json:"Bedroom"`
				Bathroom  int `json:"Bathroom"`
				Reviews   int `json:"Reviews"`
				Occupancy int `json:"Occupancy"`
			}{
				Bedroom:   3,
				Bathroom:  2,
				Reviews:   45,
				Occupancy: 6,
			},
			EcoFriendly:  true,
			FeatureImage: "https://example.com/images/feature.jpg",
			Image: &struct {
				Count  int      `json:"Count,omitempty"`
				Images []string `json:"Images,omitempty"`
			}{
				Count:  3,
				Images: []string{"image1.jpg", "image2.jpg", "image3.jpg"},
			},
			Price:                  500,
			PropertyName:           "Luxury Downtown Apartment",
			PropertySlug:           "luxury-downtown-apartment",
			PropertyType:           "Apartment",
			PropertyTypeCategoryId: "APT001",
			ReviewScore:            45,
			ReviewScores: map[string]float64{
				"Cleanliness": 4.5,
				"Location":    4.8,
				"Value":       4.2,
				"Service":     4.6,
			},
			RoomSize:  85.5,
			MinStay:   2,
			UpdatedAt: "2025-01-09T06:21:18Z",
		},
		Partner: struct {
			ID         string   `json:"ID"`
			Archived   []string `json:"Archived"`
			OwnerID    string   `json:"OwnerID"`
			HcomID     string   `json:"HcomID"`
			BrandId    string   `json:"BrandId"`
			URL        string   `json:"URL"`
			UnitNumber string   `json:"UnitNumber"`
			EpCluster  string   `json:"EpCluster"`
		}{
			ID:         "PTR123",
			Archived:   []string{"old-listing-1", "old-listing-2"},
			OwnerID:    "OWN456",
			HcomID:     "HCOM789",
			BrandId:    "BRD001",
			URL:        "https://example.com/property/luxury-downtown",
			UnitNumber: "APT-501",
			EpCluster:  "US-EAST",
		},
	}
}

// Helper function to create minimal property details response
func getMinimalPropertyDetails() structs.PropertyDetailsResponse {
	return structs.PropertyDetailsResponse{
		ID:        "PROP123",
		Feed:      1,
		Published: true,
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
			City:        "New York",
			Country:     "United States",
			CountryCode: "US",
			LocationID:  "NYC123",
		},
		Property: struct {
			Amenities map[string]string `json:"Amenities"`
			Counts    struct {
				Bedroom   int `json:"Bedroom"`
				Bathroom  int `json:"Bathroom"`
				Reviews   int `json:"Reviews"`
				Occupancy int `json:"Occupancy"`
			} `json:"Counts"`
			EcoFriendly  bool   `json:"EcoFriendly"`
			FeatureImage string `json:"FeatureImage"`
			Image        *struct {
				Count  int      `json:"Count,omitempty"`
				Images []string `json:"Images,omitempty"`
			} `json:"Image,omitempty"`
			Price                  int                `json:"Price"`
			PropertyName           string             `json:"PropertyName"`
			PropertySlug           string             `json:"PropertySlug"`
			PropertyType           string             `json:"PropertyType"`
			PropertyTypeCategoryId string             `json:"PropertyTypeCategoryId"`
			ReviewScore            int                `json:"ReviewScore"`
			ReviewScores           map[string]float64 `json:"ReviewScores,omitempty"`
			RoomSize               float64            `json:"RoomSize"`
			MinStay                int                `json:"MinStay"`
			UpdatedAt              string             `json:"UpdatedAt"`
		}{
			Price:        300,
			PropertyName: "Basic Apartment",
			PropertyType: "Apartment",
			UpdatedAt:    "2025-01-09T06:21:18Z",
		},
		Partner: struct {
			ID         string   `json:"ID"`
			Archived   []string `json:"Archived"`
			OwnerID    string   `json:"OwnerID"`
			HcomID     string   `json:"HcomID"`
			BrandId    string   `json:"BrandId"`
			URL        string   `json:"URL"`
			UnitNumber string   `json:"UnitNumber"`
			EpCluster  string   `json:"EpCluster"`
		}{
			ID:      "PTR123",
			OwnerID: "OWN456",
		},
	}
}
