package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
	web.Controller
}

func (c *PropertyDetailsController) GetPropertyDetails() {
	propertyId := c.Ctx.Input.Param(":propertyId")

	// Call external API and fetch property details
	externalAPIURL := "http://192.168.0.44:8085/dynamodb-s3-os?propertyId=" + propertyId + "&languageCode=en"
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to fetch property details")); writeErr != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			// Log the error if you have a logging mechanism
		}
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to decode response")); writeErr != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			// Log the error if you have a logging mechanism
		}
		return
	}

	c.Data["json"] = data
	if err := c.ServeJSON(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to serve JSON response")); writeErr != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			// Log the error if you have a logging mechanism
		}
	}
}
