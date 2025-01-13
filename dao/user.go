// dao/user_dao.go
package dao

import (
	"beego-api-service/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type UserDAO interface {
	Create(user *models.User) error
	GetByIdentifier(identifier string) (*models.User, error)
	Update(identifier string, user *models.User) error
	Delete(identifier string) error
}

type userDAO struct {
	orm orm.Ormer
}

func NewUserDAO() UserDAO {
	return &userDAO{
		orm: orm.NewOrm(),
	}
}

func (d *userDAO) Create(user *models.User) error {
	_, err := d.orm.Insert(user)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return fmt.Errorf("email already exists")
	}
	return err
}

// Implement other DAO methods as needed
func (d *userDAO) GetByIdentifier(identifier string) (*models.User, error) {
	user := &models.User{}
	var err error

	// Try to parse as ID first
	if id, parseErr := strconv.ParseInt(identifier, 10, 64); parseErr == nil {
		err = d.orm.QueryTable(user).Filter("id", id).One(user)
	} else {
		// Try email
		err = d.orm.QueryTable(user).Filter("email", identifier).One(user)
	}

	if err == orm.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (d *userDAO) Update(identifier string, updateData *models.User) error {
	user, err := d.GetByIdentifier(identifier)
	if err != nil {
		return err
	}

	// Update fields if provided
	if updateData.Name != "" {
		user.Name = updateData.Name
	}
	if updateData.Age != 0 {
		user.Age = updateData.Age
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}

	_, err = d.orm.Update(user)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return fmt.Errorf("email already exists")
	}
	return err
}

func (d *userDAO) Delete(identifier string) error {
	user, err := d.GetByIdentifier(identifier)
	if err != nil {
		return err
	}

	_, err = d.orm.Delete(user)
	return err
}
