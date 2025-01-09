package responses

import (
	"log"
	"net/http"

	"beego-api-service/structs"

	"github.com/beego/beego/v2/server/web"
)

func SendImagesResponse(c *web.Controller, data structs.ImagesResponse) {
	c.Data["json"] = data
	if err := c.ServeJSON(); err != nil {
		log.Printf("Failed to serve JSON response: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		if writeErr := c.Ctx.Output.Body([]byte("Failed to serve JSON response")); writeErr != nil {
			log.Printf("Failed to write error response: %v", writeErr)
		}
	}
}
