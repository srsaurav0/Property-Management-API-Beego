package responses

import "beego-api-service/structs"

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}, message string) *structs.StandardResponse {
	return &structs.StandardResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code, message string) *structs.StandardResponse {
	return &structs.StandardResponse{
		Success: false,
		Error: &structs.ErrorInfo{
			Code:    code,
			Message: message,
		},
	}
}
