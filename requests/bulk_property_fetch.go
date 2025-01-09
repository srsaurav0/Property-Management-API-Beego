package requests

import (
	"errors"
	"log"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func GetPropertyIDs(c *web.Controller) ([]string, error) {
	propertyIds := c.GetString("propertyIds")
	if propertyIds == "" {
		log.Printf("no property IDs provided")
		return nil, errors.New("no property IDs provided")
	}
	ids := strings.Split(propertyIds, ",")
	return ids, nil
}
