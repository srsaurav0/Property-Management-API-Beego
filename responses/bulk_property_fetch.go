package responses

import (
	"beego-api-service/structs"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

func SendPropertyDetailsResponses(c *web.Controller, data []structs.PropertyDetailsResponse) {
	c.Data["json"] = data
	if err := c.ServeJSON(); err != nil {
		log.Printf("Failed to serve JSON response: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to serve JSON response")); writeErr != nil {
			log.Printf("Failed to write error response: %v", writeErr)
		}
	}
}
