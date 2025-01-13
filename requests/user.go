package requests

import (
	"encoding/json"
	"fmt"
	"log"

	"beego-api-service/responses"
	"beego-api-service/validation"

	"github.com/beego/beego/v2/server/web"
)

// HandleCreateUserRequest handles the parsing and validation of create user request
func HandleCreateUserRequest(c *web.Controller, body []byte) (*validation.CreateUserRequest, interface{}) {
	var req validation.CreateUserRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("JSON Unmarshal error: %v\n", err)
		return nil, responses.NewErrorResponse("INVALID_REQUEST", fmt.Sprintf("Invalid JSON format: %v", err))
	}

	log.Printf("Successfully parsed request: %+v\n", req)

	// Validate the request
	if err := req.Validate(); err != nil {
		return nil, responses.NewErrorResponse("VALIDATION_ERROR", err.Error())
	}

	return &req, nil
}
