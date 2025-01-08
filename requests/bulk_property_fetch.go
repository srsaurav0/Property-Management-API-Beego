package requests

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func GetPropertyIDs(c *web.Controller) ([]string, error) {
	propertyIds := c.GetString("propertyIds")
	if propertyIds == "" {
		return nil, fmt.Errorf("no property IDs provided")
	}
	ids := strings.Split(propertyIds, ",")
	return ids, nil
}
