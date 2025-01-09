package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestFetchOSPropertyDetails(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request parameters
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "test123", r.URL.Query().Get("propertyId"))
		assert.Equal(t, "en", r.URL.Query().Get("languageCode"))

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"OS": map[string]interface{}{
				"id":        "test123",
				"feed":      float64(1),
				"published": true,
				"categories": `[{
					"Name": "Test Category",
					"Slug": "test-category",
					"Type": "location",
					"Display": ["Display 1", "Display 2"],
					"LocationID": "loc_123"
				}]`,
				"city":         "Test City",
				"country":      "Test Country",
				"country_code": "TC",
				"display":      "Test Display",
				"location_id":  "loc_123",
				"state_abbr":   "TS",
				"lonlat": map[string]interface{}{
					"coordinates": []interface{}{float64(67.890), float64(12.345)},
				},
				"amenity_categories": []interface{}{"WiFi", "Pool"},
				"bedroom_count":      float64(2),
				"bathroom_count":     float64(1),
				"number_of_review":   float64(10),
				"occupancy":          float64(4),
				"property_flags": map[string]interface{}{
					"eco_friendly": true,
				},
				"feature_image":          "test_image.jpg",
				"usd_price":              float64(150),
				"property_name":          "Test Property",
				"property_slug":          "test-property",
				"property_type":          "Apartment",
				"property_type_category": "apt_123",
				"review_score_general":   float64(85),
				"room_size_sqft":         float64(45.5),
				"min_stay":               float64(2),
				"updated_at":             "2024-01-09T12:00:00Z",
				"owner_id":               "owner_123",
				"hcom_id":                "hcom_123",
				"brand_id":               "brand_123",
				"feed_provider_url":      "https://example.com/property/123",
				"unit_number":            "unit_123",
				"cluster_id":             "cluster_1",
				"archived":               []interface{}{"2023-12-01"},
			},
		})
	}))
	defer mockServer.Close()

	// Set up test configuration
	web.AppConfig.Set("externalAPIBaseURL", mockServer.URL)

	tests := []struct {
		name        string
		propertyID  string
		wantErr     bool
		validateRes func(*testing.T, structs.PropertyDetailsResponse)
	}{
		{
			name:       "successful fetch",
			propertyID: "test123",
			wantErr:    false,
			validateRes: func(t *testing.T, res structs.PropertyDetailsResponse) {
				// Validate basic fields
				assert.Equal(t, "test123", res.ID)
				assert.Equal(t, 1, res.Feed)
				assert.True(t, res.Published)

				// Validate GeoInfo
				assert.Equal(t, "Test City", res.GeoInfo.City)
				assert.Equal(t, "Test Country", res.GeoInfo.Country)
				assert.Equal(t, "TC", res.GeoInfo.CountryCode)
				assert.Equal(t, "Test Display", res.GeoInfo.Display)
				assert.Equal(t, "loc_123", res.GeoInfo.LocationID)
				assert.Equal(t, "TS", res.GeoInfo.StateAbbr)
				assert.Equal(t, "12.345000", res.GeoInfo.Lat)
				assert.Equal(t, "67.890000", res.GeoInfo.Lng)

				// Validate Categories
				assert.Len(t, res.GeoInfo.Categories, 1)
				cat := res.GeoInfo.Categories[0]
				assert.Equal(t, "Test Category", cat.Name)
				assert.Equal(t, "test-category", cat.Slug)
				assert.Equal(t, "location", cat.Type)
				assert.Equal(t, []string{"Display 1", "Display 2"}, cat.Display)
				assert.Equal(t, "loc_123", cat.LocationID)

				// Validate Property
				assert.Len(t, res.Property.Amenities, 2)
				assert.Equal(t, 2, res.Property.Counts.Bedroom)
				assert.Equal(t, 1, res.Property.Counts.Bathroom)
				assert.Equal(t, 10, res.Property.Counts.Reviews)
				assert.Equal(t, 4, res.Property.Counts.Occupancy)
				assert.True(t, res.Property.EcoFriendly)
				assert.Equal(t, "test_image.jpg", res.Property.FeatureImage)
				assert.Equal(t, 150, res.Property.Price)
				assert.Equal(t, "Test Property", res.Property.PropertyName)
				assert.Equal(t, "test-property", res.Property.PropertySlug)
				assert.Equal(t, "Apartment", res.Property.PropertyType)
				assert.Equal(t, "apt_123", res.Property.PropertyTypeCategoryId)
				assert.Equal(t, 85, res.Property.ReviewScore)
				assert.Equal(t, 45.5, res.Property.RoomSize)
				assert.Equal(t, 2, res.Property.MinStay)
				assert.Equal(t, "2024-01-09T12:00:00Z", res.Property.UpdatedAt)

				// Validate Partner
				assert.Equal(t, "test123", res.Partner.ID)
				assert.Equal(t, []string{"2023-12-01"}, res.Partner.Archived)
				assert.Equal(t, "owner_123", res.Partner.OwnerID)
				assert.Equal(t, "hcom_123", res.Partner.HcomID)
				assert.Equal(t, "brand_123", res.Partner.BrandId)
				assert.Equal(t, "https://example.com/property/123", res.Partner.URL)
				assert.Equal(t, "unit_123", res.Partner.UnitNumber)
				assert.Equal(t, "cluster_1", res.Partner.EpCluster)
			},
		},
		{
			name:       "empty property ID",
			propertyID: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FetchOSPropertyDetails(tt.propertyID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.validateRes != nil {
				tt.validateRes(t, result)
			}
		})
	}
}
