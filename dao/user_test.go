package dao

import (
	"beego-api-service/models"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// testUserDAO is the test implementation
type testUserDAO struct {
	orm orm.Ormer
}

// NewTestUserDAO creates a new instance of UserDAO for testing
func NewTestUserDAO() UserDAO {
	return &testUserDAO{
		orm: orm.NewOrm(),
	}
}

func init() {
	// Set up test configuration
	err := web.LoadAppConfig("ini", "../conf/app.conf")
	if err != nil {
		panic("Failed to load app.conf for testing: " + err.Error())
	}

	// Set timestamp for tests
	web.BConfig.RunMode = "test"

	// Set any other test-specific configurations
	web.BConfig.WebConfig.AutoRender = false

	// Set test user
	web.AppConfig.Set("TestUser", "srsaurav0")

	// Register database driver
	err = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		panic(err)
	}

	// Register model
	orm.RegisterModel(new(models.User))

	// Set up the database
	err = orm.RegisterDataBase("default", "sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// Create tables
	err = orm.RunSyncdb("default", true, true)
	if err != nil {
		panic(err)
	}
}

// In the testUserDAO implementation
func (d *testUserDAO) Create(user *models.User) error {
	_, err := d.orm.Insert(user)
	if err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return fmt.Errorf("email already exists")
	}
	return err
}

func (d *testUserDAO) GetByIdentifier(identifier string) (*models.User, error) {
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

func (d *testUserDAO) Update(identifier string, updateData *models.User) error {
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

func (d *testUserDAO) Delete(identifier string) error {
	user, err := d.GetByIdentifier(identifier)
	if err != nil {
		return err
	}

	_, err = d.orm.Delete(user)
	return err
}

func TestUserDAO(t *testing.T) {
	dao := NewTestUserDAO()

	// Test Create
	t.Run("Create User", func(t *testing.T) {
		user := &models.User{
			Name:  "Test User",
			Age:   25,
			Email: "test@example.com",
		}

		err := dao.Create(user)
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)
	})

	t.Run("Create Duplicate Email", func(t *testing.T) {
		user := &models.User{
			Name:  "Another User",
			Age:   30,
			Email: "test@example.com",
		}

		err := dao.Create(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email already exists")
	})

	// Test GetByIdentifier
	t.Run("Get By Email", func(t *testing.T) {
		user, err := dao.GetByIdentifier("test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)
		assert.Equal(t, 25, user.Age)
	})

	t.Run("Get By ID", func(t *testing.T) {
		user, err := dao.GetByIdentifier("1")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)
	})

	t.Run("Get Non-existent User", func(t *testing.T) {
		user, err := dao.GetByIdentifier("nonexistent@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})

	// Test Update
	t.Run("Update User", func(t *testing.T) {
		updateUser := &models.User{
			Name: "Updated User",
			Age:  26,
		}

		err := dao.Update("test@example.com", updateUser)
		assert.NoError(t, err)

		// Verify update
		user, err := dao.GetByIdentifier("test@example.com")
		assert.NoError(t, err)
		assert.Equal(t, "Updated User", user.Name)
		assert.Equal(t, 26, user.Age)
	})

	t.Run("Update Non-existent User", func(t *testing.T) {
		updateUser := &models.User{
			Name: "Updated User",
			Age:  26,
		}

		err := dao.Update("nonexistent@example.com", updateUser)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})

	// Test Delete
	t.Run("Delete User", func(t *testing.T) {
		err := dao.Delete("test@example.com")
		assert.NoError(t, err)

		// Verify deletion
		user, err := dao.GetByIdentifier("test@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("Delete Non-existent User", func(t *testing.T) {
		err := dao.Delete("nonexistent@example.com")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

func TestNewTestUserDAO(t *testing.T) {
	dao := NewTestUserDAO()
	assert.NotNil(t, dao)

	// Assert that it implements the UserDAO interface
	_, ok := interface{}(dao).(UserDAO)
	assert.True(t, ok)
}
