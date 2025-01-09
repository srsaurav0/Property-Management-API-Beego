package services

import (
	"beego-api-service/structs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestFetchOSPropertyDetails(t *testing.T) {
	tests := []struct {
		name           string
		propertyID     string
		mockResponse   string
		expectedResult structs.PropertyDetailsResponse
		expectError    bool
	}{
		{
			name:       "Success with complete data",
			propertyID: "HA-3213808988",
			mockResponse: `{
                "OS": {
                    "amenity_categories": [
                        "Air Conditioner",
                        "Balcony/Terrace",
                        "Bedding/Linens",
                        "Child Friendly",
                        "Kitchen",
                        "Laundry",
                        "Pool",
                        "View",
                        "Ocean View",
                        "Sports/Activities",
                        "Wellness Facilities",
                        "Spa",
                        "Guest Services",
                        "Entertainment"
                    ],
                    "archived": ["VRBO", "EP", "HC"],
                    "bathroom_count": 3,
                    "bedroom_count": 3,
                    "brand_id": "321",
                    "categories": "[{\"LocationID\": \"117\", \"Name\": \"Mexico\", \"Type\": \"country\", \"Slug\": \"mexico\", \"Display\": [\"mexico\"]}, {\"LocationID\": \"11116\", \"Name\": \"Baja California Sur\", \"Type\": \"state\", \"Slug\": \"mexico/baja-california-sur\", \"Display\": [\"mexico\", \"baja-california-sur\"]}, {\"LocationID\": \"180032\", \"Name\": \"Los Cabos\", \"Type\": \"region\", \"Slug\": \"mexico/baja-california-sur/los-cabos\", \"Display\": [\"mexico\", \"baja-california-sur\", \"los-cabos\"]}, {\"LocationID\": \"6349690\", \"Name\": \"El Tezal\", \"Type\": \"city\", \"Slug\": \"mexico/baja-california-sur/cabo-san-lucas/el-tezal\", \"Display\": [\"mexico\", \"baja-california-sur\", \"cabo-san-lucas\", \"el-tezal\"]}]",
                    "city": "El Tezal",
                    "cluster_id": "c002",
                    "country": "Mexico",
                    "country_code": "MX",
                    "display": "El Tezal, Cabo San Lucas, Baja California Sur, Mexico",
                    "feature_image": "https://images.trvl-media.com/lodging/102000000/101740000/101739900/101739817/7c032e5f.jpg?impolicy=fcrop&w=1000&h=666&quality=medium",
                    "feed": 12,
                    "feed_provider_url": "https://www.vrbo.com/search?selected=101739817&regionId=6349690",
                    "hcom_id": "3256674144",
                    "id": "HA-3213808988",
                    "location_id": "6349690",
                    "lonlat": {
                        "coordinates": [-109.88175, 22.907337]
                    },
                    "min_stay": 1,
                    "number_of_review": 1,
                    "occupancy": 8,
                    "owner_id": "101739817",
                    "property_flags": {
                        "eco_friendly": false
                    },
                    "property_name": "Brand New Luxury Penthouse w/Jacuzzi",
                    "property_slug": "brand-new-luxury-penthouse-w-jacuzzi",
                    "property_type": "Apartment",
                    "property_type_category": "Apartment",
                    "published": true,
                    "review_score_general": 5,
                    "room_size_sqft": 3121,
                    "unit_number": "4383133",
                    "updated_at": "2024-05-03T11:46:19.189256+00:00",
                    "usd_price": 170
                }
            }`,
			expectedResult: structs.PropertyDetailsResponse{
				ID:        "HA-3213808988",
				Feed:      12,
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
							Name:       "Mexico",
							Slug:       "mexico",
							Type:       "country",
							Display:    []string{"mexico"},
							LocationID: "117",
						},
						{
							Name:       "Baja California Sur",
							Slug:       "mexico/baja-california-sur",
							Type:       "state",
							Display:    []string{"mexico", "baja-california-sur"},
							LocationID: "11116",
						},
						{
							Name:       "Los Cabos",
							Slug:       "mexico/baja-california-sur/los-cabos",
							Type:       "region",
							Display:    []string{"mexico", "baja-california-sur", "los-cabos"},
							LocationID: "180032",
						},
						{
							Name:       "El Tezal",
							Slug:       "mexico/baja-california-sur/cabo-san-lucas/el-tezal",
							Type:       "city",
							Display:    []string{"mexico", "baja-california-sur", "cabo-san-lucas", "el-tezal"},
							LocationID: "6349690",
						},
					},
					City:        "El Tezal",
					Country:     "Mexico",
					CountryCode: "MX",
					Display:     "El Tezal, Cabo San Lucas, Baja California Sur, Mexico",
					LocationID:  "6349690",
					StateAbbr:   "",
					Lat:         "22.907337",
					Lng:         "-109.881750",
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
						"1":  "Air Conditioner",
						"2":  "Balcony/Terrace",
						"3":  "Bedding/Linens",
						"4":  "Child Friendly",
						"5":  "Kitchen",
						"6":  "Laundry",
						"7":  "Pool",
						"8":  "View",
						"9":  "Ocean View",
						"10": "Sports/Activities",
						"11": "Wellness Facilities",
						"12": "Spa",
						"13": "Guest Services",
						"14": "Entertainment",
					},
					Counts: struct {
						Bedroom   int `json:"Bedroom"`
						Bathroom  int `json:"Bathroom"`
						Reviews   int `json:"Reviews"`
						Occupancy int `json:"Occupancy"`
					}{
						Bedroom:   3,
						Bathroom:  3,
						Reviews:   1,
						Occupancy: 8,
					},
					EcoFriendly:            false,
					FeatureImage:           "https://images.trvl-media.com/lodging/102000000/101740000/101739900/101739817/7c032e5f.jpg?impolicy=fcrop&w=1000&h=666&quality=medium",
					Price:                  170,
					PropertyName:           "Brand New Luxury Penthouse w/Jacuzzi",
					PropertySlug:           "brand-new-luxury-penthouse-w-jacuzzi",
					PropertyType:           "Apartment",
					PropertyTypeCategoryId: "Apartment",
					ReviewScore:            5,
					RoomSize:               3121,
					MinStay:                1,
					UpdatedAt:              "2024-05-03T11:46:19.189256+00:00",
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
					ID:         "HA-3213808988",
					Archived:   []string{"VRBO", "EP", "HC"},
					OwnerID:    "101739817",
					HcomID:     "3256674144",
					BrandId:    "321",
					URL:        "https://www.vrbo.com/search?selected=101739817&regionId=6349690",
					UnitNumber: "4383133",
					EpCluster:  "c002",
				},
			},
			expectError: false,
		},
		{
			name:       "Invalid OS data structure",
			propertyID: "INVALID-ID",
			mockResponse: `{
                "OS": "invalid"
            }`,
			expectedResult: structs.PropertyDetailsResponse{},
			expectError:    true,
		},
		{
			name:         "Invalid JSON response",
			propertyID:   "ERROR-ID",
			mockResponse: `invalid json`,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "GET", r.Method)
				assert.Contains(t, r.URL.String(), tt.propertyID)
				assert.Contains(t, r.URL.String(), "languageCode=en")

				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			web.AppConfig.Set("externalAPIBaseURL", server.URL)

			result, err := FetchOSPropertyDetails(tt.propertyID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestFetchOSPropertyDetailsHTTPError(t *testing.T) {
	web.AppConfig.Set("externalAPIBaseURL", "http://invalid-url-that-will-fail")

	result, err := FetchOSPropertyDetails("PROP123")

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestFetchOSPropertyDetailsConfigError(t *testing.T) {
	web.AppConfig.Set("externalAPIBaseURL", "")

	result, err := FetchOSPropertyDetails("PROP123")

	assert.Error(t, err)
	assert.Empty(t, result)
}
