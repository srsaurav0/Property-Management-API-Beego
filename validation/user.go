package validation

import "fmt"

type CreateUserRequest struct {
	// @Description The name of the user
	// @Required
	Name string `json:"name" valid:"Required"`
	// @Description The age of the user
	// @Required
	// @Minimum 0
	Age int `json:"age" valid:"Required"`
	// @Description The email address of the user
	// @Required
	// @Pattern [a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}
	Email string `json:"email" valid:"Required;Email"`
}

type UpdateUserRequest struct {
	// @Description The name of the user
	// @Required
	Name string `json:"name,omitempty"`
	// @Description The age of the user
	// @Required
	// @Minimum 0
	Age int `json:"age,omitempty"`
}

// Validate CreateUserRequest
func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Age <= 0 {
		return fmt.Errorf("invalid age")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

// Validate UpdateUserRequest
func (r *UpdateUserRequest) Validate() error {
	// If both name and age are empty, nothing to update
	if r.Name == "" && r.Age == 0 {
		return fmt.Errorf("at least one field (name or age) must be provided for update")
	}

	// Validate age if provided
	if r.Age < 0 {
		return fmt.Errorf("invalid age")
	}

	return nil
}
