package model

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error is occurred  on .env file please check")
	}
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if err != nil {
		return err
	} else {
		DB = db
		slog.Info("Successfully Connected to Database")
	}
	return nil
}

func RunMigrations() error {
	slog.Info("Running Migrations for the models.")
	return DB.AutoMigrate(
		&CustomerData{},
		&User{},
		&UserRole{},
		&Project{},
		&Organization{},
		&Role{},
		&Permission{},
		&RolePermission{},
	)
}
