package services

import (
	"beego-api-service/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

// const externalAPIBaseURL = "http://192.168.0.44:8085/dynamodb-s3-os"

func FetchPropertyImages(propertyId string) (structs.ImagesResponse, error) {
	transformedData := make(structs.ImagesResponse)

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

	galleryData, ok := originalData["S3-Gallery"].(map[string]interface{})
	if !ok {
		return transformedData, fmt.Errorf("invalid S3-Gallery format")
	}

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
