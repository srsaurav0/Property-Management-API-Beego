package responses

import (
	"beego-api-service/structs"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

func SendPropertyDetailsResponse(c *web.Controller, data structs.PropertyDetailsResponse) {
	c.Data["json"] = data
	if err := c.ServeJSON(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to serve JSON response")); writeErr != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			log.Fatalf("Failed to serve JSON response")
		}
	}
}

func SendErrorResponse(c *web.Controller, message string, status int) {
	c.Ctx.Output.SetStatus(status)
	if writeErr := c.Ctx.Output.Body([]byte(message)); writeErr != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		log.Fatalf("Failed to write error response")
	}
}
