package models

// swagger:model
type SwaggerUser struct {
	// User ID
	// example: 1
	ID int64 `json:"id"`

	// User's full name
	// example: John Doe
	Name string `json:"name"`

	// User's age
	// example: 25
	Age int `json:"age"`

	// User's email address
	// example: john.doe@example.com
	Email string `json:"email"`
}
