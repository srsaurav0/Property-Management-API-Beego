package structs

// UserCreateRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" valid:"Required"`
	Age   int    `json:"age" valid:"Required"`
	Email string `json:"email" valid:"Required;Email"`
}

// UserUpdateRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// Response represents the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error details in the response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StandardResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
