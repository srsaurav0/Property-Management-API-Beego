package controllers

import (
	"beego-api-service/responses"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

// HandleError handles common error responses
func (b *BaseController) HandleError(statusCode int, message string, err error) {
	if err != nil {
		log.Printf("[ERROR] %s: %v", message, err)
	}
	b.Data["json"] = responses.NewErrorResponse(http.StatusText(statusCode), message)
	b.Ctx.Output.SetStatus(statusCode)
	b.ServeJSON()
}

// Common error handling methods
func (c *BaseController) HandleBadRequest(message string, err error) {
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"message": message,
		"error":   err.Error(),
	}
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.ServeJSON()
}

func (c *BaseController) HandleInternalServerError(message string, err error) {
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"message": message,
		"error":   err.Error(),
	}
	c.Ctx.Output.SetStatus(http.StatusInternalServerError)
	c.ServeJSON()
}

func (b *BaseController) HandleUnauthorized(message string, err error) {
	b.HandleError(http.StatusUnauthorized, message, err)
}

func (b *BaseController) HandleForbidden(message string, err error) {
	b.HandleError(http.StatusForbidden, message, err)
}

func (c *BaseController) HandleNotFound(message string, err error) {
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"message": message,
		"error":   err.Error(),
	}
	c.Ctx.Output.SetStatus(http.StatusNotFound)
	c.ServeJSON()
}

// HandleSuccess handles successful responses
func (b *BaseController) HandleSuccess(statusCode int, data interface{}, message string) {
	log.Printf("[INFO] Success: %s", message)
	b.Data["json"] = responses.NewSuccessResponse(data, message)
	b.Ctx.Output.SetStatus(statusCode)
	b.ServeJSON()
}
