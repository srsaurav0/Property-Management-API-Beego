// @Title User API
// @Description User management API
// @Version 1.0
// @Contact.email support@example.com
// @BasePath /v1/api/user
package controllers

import (
	"beego-api-service/models"
	"beego-api-service/services"
	"beego-api-service/validation"
	"encoding/json"
	"io"
	"net/http"
)

type UserController struct {
	BaseController
	userService services.UserService
}

func (c *UserController) Prepare() {
	c.userService = services.NewUserService()
}

// Additional controller annotations for CreateUser method:
// @Title Create User
// @Description Create a new user
// @Param   body        body    validation.CreateUserRequest  true  "User information"
// @Success 201 {object} responses.StandardResponse "User created successfully"
// @Failure 400 {object} responses.StandardResponse "Bad Request"
// @Failure 500 {object} responses.StandardResponse "Internal Server Error"
// @Router /v1/api/user [post]
func (c *UserController) CreateUser() {
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

// @Title Get User
// @Description Get user by ID or email
// @Param   identifier  path    string  true  "User ID or email" example("1")
// @Success 200 {object} responses.StandardResponse "User retrieved successfully"
// @Failure 404 {object} responses.UserNotFoundResponse "User not found"
// @Failure 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /v1/api/user/{identifier} [get]
func (c *UserController) GetUser() {
	identifier := c.Ctx.Input.Param(":identifier")
	if identifier == "" {
		c.HandleBadRequest("Identifier is required", nil)
		return
	}

	user, err := c.userService.GetByIdentifier(identifier)
	if err != nil {
		if err.Error() == "user not found" {
			c.HandleNotFound("User not found", err)
			return
		}
		c.HandleInternalServerError("Failed to get user", err)
		return
	}

	c.HandleSuccess(http.StatusOK, user, "User retrieved successfully")
}

// @Title Update User
// @Description Update user by ID or email
// @Param   identifier  path    string  true  "User ID or email"
// @Param   body        body    validation.UpdateUserRequest  true  "User information to update"
// @Success 200 {object} responses.StandardResponse "User updated successfully"
// @Failure 400 {object} responses.StandardResponse "Bad Request"
// @Failure 404 {object} responses.UserNotFoundResponse "User not found"
// @Failure 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /v1/api/user/{identifier} [put]
func (c *UserController) UpdateUser() {
	identifier := c.Ctx.Input.Param(":identifier")
	if identifier == "" {
		c.HandleBadRequest("Identifier is required", nil)
		return
	}

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

	var req validation.UpdateUserRequest
	if err := json.Unmarshal(body, &req); err != nil {
		c.HandleBadRequest("Invalid JSON format", err)
		return
	}

	if err := req.Validate(); err != nil {
		c.HandleBadRequest(err.Error(), nil)
		return
	}

	user := &models.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := c.userService.Update(identifier, user); err != nil {
		if err.Error() == "user not found" {
			c.HandleNotFound("User not found", err)
			return
		}
		c.HandleInternalServerError("Failed to update user", err)
		return
	}

	updatedUser, err := c.userService.GetByIdentifier(identifier)
	if err != nil {
		c.HandleInternalServerError("Failed to fetch updated user", err)
		return
	}

	c.HandleSuccess(http.StatusOK, updatedUser, "User updated successfully")
}

// @Title Delete User
// @Description Delete user by ID or email
// @Param   identifier  path    string  true  "User ID or email"
// @Success 200 {object} responses.StandardResponse "User deleted successfully"
// @Failure 404 {object} responses.UserNotFoundResponse "User not found"
// @Failure 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /v1/api/user/{identifier} [delete]
func (c *UserController) DeleteUser() {
	identifier := c.Ctx.Input.Param(":identifier")
	if identifier == "" {
		c.HandleBadRequest("Identifier is required", nil)
		return
	}

	if err := c.userService.Delete(identifier); err != nil {
		if err.Error() == "user not found" {
			c.HandleNotFound("User not found", err)
			return
		}
		c.HandleInternalServerError("Failed to delete user", err)
		return
	}

	c.HandleSuccess(http.StatusOK, nil, "User deleted successfully")
}
