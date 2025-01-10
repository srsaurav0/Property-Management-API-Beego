package validation

import "fmt"

type CreateUserRequest struct {
	Name  string `json:"name" valid:"Required"`
	Age   int    `json:"age" valid:"Required"`
	Email string `json:"email" valid:"Required;Email"`
}

type UpdateUserRequest struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
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
