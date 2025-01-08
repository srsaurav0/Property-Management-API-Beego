package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"fmt"
)

// const externalAPIBaseURL = "http://192.168.0.44:8085/dynamodb-s3-os"

// func FetchPropertyDetails(propertyId string) (structs.PropertyDetailsResponse, error) {
// 	var transformedData structs.PropertyDetailsResponse

// 	externalAPIURL := fmt.Sprintf("%s?propertyId=%s&languageCode=en", externalAPIBaseURL, propertyId)
// 	resp, err := http.Get(externalAPIURL)
// 	if err != nil {
// 		return transformedData, fmt.Errorf("HTTP request failed: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	var originalData map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&originalData); err != nil {
// 		return transformedData, fmt.Errorf("failed to decode response: %w", err)
// 	}

// 	if err := transformData(originalData, &transformedData); err != nil {
// 		return transformedData, fmt.Errorf("failed to transform data: %w", err)
// 	}

// 	return transformedData, nil
// }

func transformOSData(originalData map[string]interface{}, transformedData *structs.PropertyDetailsResponse) error {
	osData, ok := originalData["OS"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid OS data")
	}

	transformedData.ID = osData["id"].(string)
	transformedData.Feed = int(osData["feed"].(float64))
	transformedData.Published = osData["published"].(bool)

	if err := parseCategories(osData, transformedData); err != nil {
		return err
	}

	transformedData.GeoInfo.City = osData["city"].(string)
	transformedData.GeoInfo.Country = osData["country"].(string)
	transformedData.GeoInfo.CountryCode = osData["country_code"].(string)
	transformedData.GeoInfo.Display = osData["display"].(string)
	transformedData.GeoInfo.LocationID = osData["location_id"].(string)
	transformedData.GeoInfo.StateAbbr = osData["state_abbr"].(string)
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

	return nil
}

func parseCategories(osData map[string]interface{}, transformedData *structs.PropertyDetailsResponse) error {
	if categoriesJSON, ok := osData["categories"].(string); ok {
		var categories []map[string]interface{}
		if err := json.Unmarshal([]byte(categoriesJSON), &categories); err != nil {
			return fmt.Errorf("failed to unmarshal categories: %w", err)
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
	return nil
}
