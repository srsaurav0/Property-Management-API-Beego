package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

func FetchPropertyDetails(propertyId string) (structs.PropertyDetailsResponse, error) {
	var transformedData structs.PropertyDetailsResponse

	externalAPIBaseURL, err := web.AppConfig.String("externalAPIBaseURL")
	if err != nil {
		log.Printf("failed to load external API URL from config: %v", err)
		return transformedData, err
	}

	externalAPIURL := fmt.Sprintf("%s?propertyId=%s&languageCode=en", externalAPIBaseURL, propertyId)
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return transformedData, err
	}
	defer resp.Body.Close()

	var originalData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&originalData); err != nil {
		log.Printf("failed to decode response: %v", err)
		return transformedData, err
	}

	if err := transformData(originalData, &transformedData); err != nil {
		log.Printf("failed to transform data: %v", err)
		return transformedData, err
	}

	return transformedData, nil
}

func transformData(originalData map[string]interface{}, transformedData *structs.PropertyDetailsResponse) error {
	s3Data, ok := originalData["S3"].(map[string]interface{})
	if !ok {
		log.Printf("invalid S3 data: expected map[string]interface{}, got %T", originalData["S3"])
		return errors.New("invalid S3 data")
	}

	transformedData.ID = s3Data["ID"].(string)
	transformedData.Feed = int(s3Data["Feed"].(float64))
	transformedData.Published = s3Data["Published"].(bool)

	geoInfo := s3Data["GeoInfo"].(map[string]interface{})
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

	property := s3Data["Property"].(map[string]interface{})
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

	if image, ok := property["Image"].(map[string]interface{}); ok {
		if len(image) > 0 {
			imgStruct := &struct {
				Count  int      `json:"Count,omitempty"`
				Images []string `json:"Images,omitempty"`
			}{
				Count:  int(image["Count"].(float64)),
				Images: make([]string, len(image["Images"].([]interface{}))),
			}
			for i, img := range image["Images"].([]interface{}) {
				imgStruct.Images[i] = img.(string)
			}
			transformedData.Property.Image = imgStruct
		} else {
			transformedData.Property.Image = nil
		}
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

	partner := s3Data["Partner"].(map[string]interface{})
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
