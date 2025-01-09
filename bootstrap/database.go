package bootstrap

import (
	"fmt"
	"log/slog"

	"github.com/krishnatrea/cdp/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *Config) (*gorm.DB, error) {

	host := config.PostgresHost
	port := config.DBPort
	user := config.PostgresUser
	dbname := config.PostgresDB
	pass := config.PostgresPassword

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	db, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	slog.Info("Running Migrations for the models.")
	return db.AutoMigrate(
		&model.CustomerData{},
		&model.User{},
		&model.UserRole{},
		&model.Project{},
		&model.Organization{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
	)
}
