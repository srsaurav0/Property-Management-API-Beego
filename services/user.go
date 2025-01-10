// services/user_service.go
package services

import (
	"beego-api-service/dao"
	"beego-api-service/models"
)

type UserService interface {
	Post(user *models.User) error
	GetByIdentifier(identifier string) (*models.User, error)
	Update(identifier string, user *models.User) error
	Delete(identifier string) error
}

type userService struct {
	userDAO dao.UserDAO
}

func NewUserService() UserService {
	return &userService{
		userDAO: dao.NewUserDAO(),
	}
}

func (s *userService) Post(user *models.User) error {
	return s.userDAO.Create(user)
}

func (s *userService) GetByIdentifier(identifier string) (*models.User, error) {
	return s.userDAO.GetByIdentifier(identifier)
}

func (s *userService) Update(identifier string, user *models.User) error {
	return s.userDAO.Update(identifier, user)
}

func (s *userService) Delete(identifier string) error {
	return s.userDAO.Delete(identifier)
}
