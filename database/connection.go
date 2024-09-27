package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	LoadEnv()
	fmt.Println("ok")

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		return fmt.Errorf("missing database configuration")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	var err error
	if DB, err = sql.Open("mysql", dsn); err != nil {
		return fmt.Errorf("failed connecting database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	fmt.Println("ok")
	return nil
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading environment: %v", err)
	}
}
