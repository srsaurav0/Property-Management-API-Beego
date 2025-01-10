// controllers/user.go
package controllers

import (
	"beego-api-service/models"
	"beego-api-service/services"
	"beego-api-service/validation"
	"encoding/json"
	"io"
	"net/http"
)

type CreateController struct {
	BaseController
	userService services.UserService
}

func (c *CreateController) Prepare() {
	c.userService = services.NewUserService()
}

func (c *CreateController) CreateUser() {
	// Read the body using ioutil
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.HandleBadRequest("Failed to read request body", err)
		return
	}
	defer c.Ctx.Request.Body.Close()

	if len(body) == 0 {
		c.HandleBadRequest("Request body is empty", nil)
		return
	}

	var req validation.CreateUserRequest
	if err := json.Unmarshal(body, &req); err != nil {
		c.HandleBadRequest("Invalid JSON format", err)
		return
	}

	if err := req.Validate(); err != nil {
		c.HandleBadRequest(err.Error(), nil)
		return
	}

	user := &models.User{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	if err := c.userService.Post(user); err != nil {
		c.HandleInternalServerError("Failed to create user", err)
		return
	}

	c.HandleSuccess(http.StatusCreated, user, "User created successfully")
}
