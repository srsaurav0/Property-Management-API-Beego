package requests

import (
	"errors"
	"log"

	"github.com/beego/beego/v2/server/web"
)

func GetPropertyID(c *web.Controller) (string, error) {
	propertyId := c.Ctx.Input.Param(":propertyId")
	if propertyId == "" {
		log.Printf("property ID not provided")
		return "", errors.New("property ID not provided")
	}
	return propertyId, nil
}
