package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/beego/beego/v2/server/web"
)

type BulkPropertyFetchController struct {
	web.Controller
}

func (c *BulkPropertyFetchController) BulkPropertyFetch() {
	propertyIds := c.GetString("propertyIds")
	ids := strings.Split(propertyIds, ",")

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]PropertyDetailsResponse, len(ids))

	for i, id := range ids {
		wg.Add(1)
		go func(i int, id string) {
			defer wg.Done()
			externalAPIURL := "http://192.168.0.44:8085/dynamodb-s3-os?propertyId=" + id + "&languageCode=en"
			resp, err := http.Get(externalAPIURL)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			var originalData map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&originalData); err != nil {
				return
			}

			osData, ok := originalData["OS"].(map[string]interface{})
			if !ok {
				return
			}

			transformedData := PropertyDetailsResponse{}

			if id, ok := osData["id"].(string); ok {
				transformedData.ID = id
			}
			if feed, ok := osData["feed"].(float64); ok {
				transformedData.Feed = int(feed)
			}
			if published, ok := osData["published"].(bool); ok {
				transformedData.Published = published
			}

			if categoriesJSON, ok := osData["categories"].(string); ok {
				var categories []map[string]interface{}
				if err := json.Unmarshal([]byte(categoriesJSON), &categories); err != nil {
					return
				}
				for _, category := range categories {
					transformedData.GeoInfo.Categories = append(transformedData.GeoInfo.Categories, struct {
						Name       string   `json:"Name"`
						Slug       string   `json:"Slug"`
						Type       string   `json:"Type"`
						Display    []string `json:"Display"`
						LocationID string   `json:"LocationID"`
					}{
						Name: category["Name"].(string),
						Slug: category["Slug"].(string),
						Type: category["Type"].(string),
						Display: func() []string {
							display := []string{}
							if d, ok := category["Display"].([]interface{}); ok {
								for _, v := range d {
									display = append(display, v.(string))
								}
							}
							return display
						}(),
						LocationID: category["LocationID"].(string),
					})
				}
			}

			if city, ok := osData["city"].(string); ok {
				transformedData.GeoInfo.City = city
			}
			if country, ok := osData["country"].(string); ok {
				transformedData.GeoInfo.Country = country
			}
			if countryCode, ok := osData["country_code"].(string); ok {
				transformedData.GeoInfo.CountryCode = countryCode
			}
			if display, ok := osData["display"].(string); ok {
				transformedData.GeoInfo.Display = display
			}
			if locationID, ok := osData["location_id"].(string); ok {
				transformedData.GeoInfo.LocationID = locationID
			}
			if stateAbbr, ok := osData["state_abbr"].(string); ok {
				transformedData.GeoInfo.StateAbbr = stateAbbr
			}
			if lonlat, ok := osData["lonlat"].(map[string]interface{}); ok {
				if coordinates, ok := lonlat["coordinates"].([]interface{}); ok && len(coordinates) >= 2 {
					transformedData.GeoInfo.Lat = fmt.Sprintf("%f", coordinates[1].(float64))
					transformedData.GeoInfo.Lng = fmt.Sprintf("%f", coordinates[0].(float64))
				}
			}

			transformedData.Property.Amenities = func() map[string]string {
				amenities := map[string]string{}
				if amenitiesList, ok := osData["amenity_categories"].([]interface{}); ok {
					for i, amenity := range amenitiesList {
						amenities[fmt.Sprintf("%d", i+1)] = amenity.(string)
					}
				}
				return amenities
			}()

			if bedroomCount, ok := osData["bedroom_count"].(float64); ok {
				transformedData.Property.Counts.Bedroom = int(bedroomCount)
			}
			if bathroomCount, ok := osData["bathroom_count"].(float64); ok {
				transformedData.Property.Counts.Bathroom = int(bathroomCount)
			}
			if numberOfReview, ok := osData["number_of_review"].(float64); ok {
				transformedData.Property.Counts.Reviews = int(numberOfReview)
			}
			if occupancy, ok := osData["occupancy"].(float64); ok {
				transformedData.Property.Counts.Occupancy = int(occupancy)
			}

			if propertyFlags, ok := osData["property_flags"].(map[string]interface{}); ok {
				if ecoFriendly, ok := propertyFlags["eco_friendly"].(bool); ok {
					transformedData.Property.EcoFriendly = ecoFriendly
				}
			}
			if featureImage, ok := osData["feature_image"].(string); ok {
				transformedData.Property.FeatureImage = featureImage
			}

			if usdPrice, ok := osData["usd_price"].(float64); ok {
				transformedData.Property.Price = int(usdPrice)
			}
			if propertyName, ok := osData["property_name"].(string); ok {
				transformedData.Property.PropertyName = propertyName
			}
			if propertySlug, ok := osData["property_slug"].(string); ok {
				transformedData.Property.PropertySlug = propertySlug
			}
			if propertyType, ok := osData["property_type"].(string); ok {
				transformedData.Property.PropertyType = propertyType
			}
			if propertyTypeCategory, ok := osData["property_type_category"].(string); ok {
				transformedData.Property.PropertyTypeCategoryId = propertyTypeCategory
			}
			if reviewScoreGeneral, ok := osData["review_score_general"].(float64); ok {
				score := int(reviewScoreGeneral)
				transformedData.Property.ReviewScore = score
			}
			if reviewScores, ok := osData["review_scores"].(map[string]interface{}); ok {
				transformedData.Property.ReviewScores = make(map[string]float64)
				for k, v := range reviewScores {
					if score, ok := v.(float64); ok {
						transformedData.Property.ReviewScores[k] = score
					}
				}
			}
			if roomSizeSqft, ok := osData["room_size_sqft"].(float64); ok {
				transformedData.Property.RoomSize = roomSizeSqft
			}
			if minStay, ok := osData["min_stay"].(float64); ok {
				transformedData.Property.MinStay = int(minStay)
			}
			if updatedAt, ok := osData["updated_at"].(string); ok {
				transformedData.Property.UpdatedAt = updatedAt
			}

			if partnerID, ok := osData["id"].(string); ok {
				transformedData.Partner.ID = partnerID
			}
			if archived, ok := osData["archived"].([]interface{}); ok {
				for _, arch := range archived {
					if archStr, ok := arch.(string); ok {
						transformedData.Partner.Archived = append(transformedData.Partner.Archived, archStr)
					}
				}
			}
			if ownerID, ok := osData["owner_id"].(string); ok {
				transformedData.Partner.OwnerID = ownerID
			}
			if hcomID, ok := osData["hcom_id"].(string); ok {
				transformedData.Partner.HcomID = hcomID
			}
			if brandId, ok := osData["brand_id"].(string); ok {
				transformedData.Partner.BrandId = brandId
			}
			if feedProviderURL, ok := osData["feed_provider_url"].(string); ok {
				transformedData.Partner.URL = feedProviderURL
			}
			if unitNumber, ok := osData["unit_number"].(string); ok {
				transformedData.Partner.UnitNumber = unitNumber
			}
			if clusterID, ok := osData["cluster_id"].(string); ok {
				transformedData.Partner.EpCluster = clusterID
			}

			mu.Lock()
			results[i] = transformedData
			mu.Unlock()
		}(i, id)
	}
	wg.Wait()

	c.Data["json"] = results
	if err := c.ServeJSON(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to serve JSON response")); writeErr != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			// Log the error if you have a logging mechanism
		}
	}
}
