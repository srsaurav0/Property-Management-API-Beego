package controllers

import (
	"log"
	"net/http"

	"beego-api-service/requests"
	"beego-api-service/responses"
	"beego-api-service/services"

	"github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
	web.Controller
}

func (c *PropertyDetailsController) GetPropertyDetails() {
	propertyId, err := requests.GetPropertyID(&c.Controller)
	if err != nil {
		log.Println(err)
		responses.SendErrorResponse(&c.Controller, "Property ID not provided", http.StatusBadRequest)
		return
	}

	transformedData, err := services.FetchPropertyDetails(propertyId)
	if err != nil {
		log.Println(err)
		responses.SendErrorResponse(&c.Controller, "Failed to fetch property details", http.StatusInternalServerError)
		return
	}

	responses.SendPropertyDetailsResponse(&c.Controller, transformedData)
}
