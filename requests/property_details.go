package requests

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func GetPropertyID(c *web.Controller) (string, error) {
	propertyId := c.Ctx.Input.Param(":propertyId")
	if propertyId == "" {
		return "", fmt.Errorf("property ID not provided")
	}
	return propertyId, nil
}
