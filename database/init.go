package database

import (
	"beego-api-service/models"
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func Init() {
	// Register database driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Log the connection details (remove in production)
	logs.Info("Connecting to database with following details:")
	logs.Info("Host:", dbHost)
	logs.Info("Port:", dbPort)
	logs.Info("Database:", dbName)
	logs.Info("User:", dbUser)

	// Construct the connection string
	datasource := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// Register default database
	err := orm.RegisterDataBase("default", "postgres", datasource)
	if err != nil {
		logs.Error("Failed to register database:", err)
		return
	}

	// Register model
	orm.RegisterModel(
		new(models.User),
	)

	// Create tables
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("Failed to sync database:", err)
		return
	}

	logs.Info("Database initialized successfully")
}
