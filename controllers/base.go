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
func (b *BaseController) HandleBadRequest(message string, err error) {
	b.HandleError(http.StatusBadRequest, message, err)
}

func (b *BaseController) HandleInternalServerError(message string, err error) {
	b.HandleError(http.StatusInternalServerError, message, err)
}

func (b *BaseController) HandleUnauthorized(message string, err error) {
	b.HandleError(http.StatusUnauthorized, message, err)
}

func (b *BaseController) HandleForbidden(message string, err error) {
	b.HandleError(http.StatusForbidden, message, err)
}

func (b *BaseController) HandleNotFound(message string, err error) {
	b.HandleError(http.StatusNotFound, message, err)
}

// HandleSuccess handles successful responses
func (b *BaseController) HandleSuccess(statusCode int, data interface{}, message string) {
	log.Printf("[INFO] Success: %s", message)
	b.Data["json"] = responses.NewSuccessResponse(data, message)
	b.Ctx.Output.SetStatus(statusCode)
	b.ServeJSON()
}
