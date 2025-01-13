package responses

// StandardResponse represents the standard API response structure
// @Description Standard API response format
type StandardResponse struct {
	// @Description HTTP status code
	// @Example 200
	Code int `json:"code" example:"200"`

	// @Description Response message
	// @Example "Operation successful"
	Message string `json:"message" example:"Operation successful"`

	// @Description Response data
	Data interface{} `json:"data"`
}

type UserNotFoundResponse struct {
	// @Description HTTP status code
	// @Example 404
	Code int `json:"code" example:"404"`

	// @Description Response message
	// @Example "Operation failed"
	Message string `json:"message" example:"Operation failed"`

	// @Description Response data
	Data interface{} `json:"data"`
}

type InternalServerErrorResponse struct {
	// @Description HTTP status code
	// @Example 500
	Code int `json:"code" example:"500"`

	// @Description Response message
	// @Example "Internal Server Error"
	Message string `json:"message" example:"Operation failed"`

	// @Description Response data
	Data interface{} `json:"data"`
}
