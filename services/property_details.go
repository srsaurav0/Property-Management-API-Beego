package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

func FetchPropertyDetails(propertyId string) (structs.PropertyDetailsResponse, error) {
	var transformedData structs.PropertyDetailsResponse

	externalAPIBaseURL, err := web.AppConfig.String("externalAPIBaseURL")
	if err != nil {
		return transformedData, fmt.Errorf("failed to load external API URL from config: %w", err)
	}

	externalAPIURL := fmt.Sprintf("%s?propertyId=%s&languageCode=en", externalAPIBaseURL, propertyId)
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		return transformedData, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var originalData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&originalData); err != nil {
		return transformedData, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := transformData(originalData, &transformedData); err != nil {
		return transformedData, fmt.Errorf("failed to transform data: %w", err)
	}

	return transformedData, nil
}

func transformData(originalData map[string]interface{}, transformedData *structs.PropertyDetailsResponse) error {
	osData, ok := originalData["S3"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid OS data")
	}

	transformedData.ID = osData["ID"].(string)
	transformedData.Feed = int(osData["Feed"].(float64))
	transformedData.Published = osData["Published"].(bool)

	geoInfo := osData["GeoInfo"].(map[string]interface{})
	for _, category := range geoInfo["Categories"].([]interface{}) {
		cat := category.(map[string]interface{})
		transformedData.GeoInfo.Categories = append(transformedData.GeoInfo.Categories, struct {
			Name       string   `json:"Name"`
			Slug       string   `json:"Slug"`
			Type       string   `json:"Type"`
			Display    []string `json:"Display"`
			LocationID string   `json:"LocationID"`
		}{
			Name: cat["Name"].(string),
			Slug: cat["Slug"].(string),
			Type: cat["Type"].(string),
			Display: func() []string {
				display := []string{}
				for _, d := range cat["Display"].([]interface{}) {
					display = append(display, d.(string))
				}
				return display
			}(),
			LocationID: cat["LocationID"].(string),
		})
	}

	transformedData.GeoInfo.City = geoInfo["City"].(string)
	transformedData.GeoInfo.Country = geoInfo["Country"].(string)
	transformedData.GeoInfo.CountryCode = geoInfo["CountryCode"].(string)
	transformedData.GeoInfo.Display = geoInfo["Display"].(string)
	transformedData.GeoInfo.LocationID = geoInfo["LocationID"].(string)
	transformedData.GeoInfo.StateAbbr = geoInfo["StateAbbr"].(string)
	transformedData.GeoInfo.Lat = geoInfo["Lat"].(string)
	transformedData.GeoInfo.Lng = geoInfo["Lng"].(string)

	property := osData["Property"].(map[string]interface{})
	transformedData.Property.Amenities = func() map[string]string {
		amenities := map[string]string{}
		for k, v := range property["Amenities"].(map[string]interface{}) {
			amenities[k] = v.(string)
		}
		return amenities
	}()
	counts := property["Counts"].(map[string]interface{})
	transformedData.Property.Counts.Bedroom = int(counts["Bedroom"].(float64))
	transformedData.Property.Counts.Bathroom = int(counts["Bathroom"].(float64))
	transformedData.Property.Counts.Reviews = int(counts["Reviews"].(float64))
	transformedData.Property.Counts.Occupancy = int(counts["Occupancy"].(float64))

	transformedData.Property.EcoFriendly = property["EcoFriendly"].(bool)
	transformedData.Property.FeatureImage = property["FeatureImage"].(string)
	image := property["Image"].(map[string]interface{})
	transformedData.Property.Image.Count = int(image["Count"].(float64))
	for _, img := range image["Images"].([]interface{}) {
		transformedData.Property.Image.Images = append(transformedData.Property.Image.Images, img.(string))
	}

	transformedData.Property.Price = int(property["Price"].(float64))
	transformedData.Property.PropertyName = property["PropertyName"].(string)
	transformedData.Property.PropertySlug = property["PropertySlug"].(string)
	transformedData.Property.PropertyType = property["PropertyType"].(string)
	transformedData.Property.PropertyTypeCategoryId = property["PropertyTypeCategoryId"].(string)
	transformedData.Property.ReviewScore = int(property["ReviewScore"].(float64))
	transformedData.Property.ReviewScores = func() map[string]float64 {
		reviewScores := map[string]float64{}
		for k, v := range property["ReviewScores"].(map[string]interface{}) {
			reviewScores[k] = v.(float64)
		}
		return reviewScores
	}()
	transformedData.Property.RoomSize = property["RoomSize"].(float64)
	transformedData.Property.MinStay = int(property["MinStay"].(float64))
	transformedData.Property.UpdatedAt = property["UpdatedAt"].(string)

	partner := osData["Partner"].(map[string]interface{})
	transformedData.Partner.ID = partner["ID"].(string)
	for _, archived := range partner["Archived"].([]interface{}) {
		transformedData.Partner.Archived = append(transformedData.Partner.Archived, archived.(string))
	}
	transformedData.Partner.OwnerID = partner["OwnerID"].(string)
	transformedData.Partner.HcomID = partner["HcomID"].(string)
	transformedData.Partner.BrandId = partner["BrandId"].(string)
	transformedData.Partner.URL = partner["URL"].(string)
	transformedData.Partner.UnitNumber = partner["UnitNumber"].(string)
	transformedData.Partner.EpCluster = partner["EpCluster"].(string)

	return nil
}
