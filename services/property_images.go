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

func FetchPropertyImages(propertyId string) (structs.ImagesResponse, error) {
	transformedData := make(structs.ImagesResponse)

	// Load the external API base URL from the configuration
	externalAPIBaseURL, err := web.AppConfig.String("externalAPIBaseURL")
	if err != nil {
		log.Printf("failed to load external API URL from config: %v", err)
		return transformedData, err
	}

	// Construct the external API URL
	externalAPIURL := fmt.Sprintf("%s?propertyId=%s&languageCode=en", externalAPIBaseURL, propertyId)
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return transformedData, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var originalData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&originalData); err != nil {
		log.Printf("failed to decode response: %v", err)
		return transformedData, err
	}

	// Extract gallery data
	galleryData, ok := originalData["S3-Gallery"].(map[string]interface{})
	if !ok {
		log.Printf("invalid S3-Gallery format: expected map[string]interface{}, got %T", originalData["S3-Gallery"])
		return transformedData, errors.New("invalid S3-Gallery format")
	}

	// Transform the gallery data
	for _, images := range galleryData {
		for _, image := range images.([]interface{}) {
			img := image.(map[string]interface{})
			label := img["label"].(string)
			url := img["url"].(string)

			if _, ok := transformedData[label]; !ok {
				transformedData[label] = []string{}
			}
			transformedData[label] = append(transformedData[label], url)
		}
	}

	return transformedData, nil
}
