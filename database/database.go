package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {
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
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		slog.Error("There is an error while connecting to the database", "Error", errSql)
	} else {
		Db = db
		slog.Info("Successfully connected to database!")
	}
}
