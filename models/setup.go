package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DbPassword := os.Getenv("DB_PASSWORD")

	// MySQL DSN 格式
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open("mysql", db)
	if err != nil {
		fmt.Println("Cannot connect to database ", DbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", DbDriver)
	}

	DB.AutoMigrate(&User{})
}

// Cleanup, db connect closed after exits
func Cleanup() {
	if DB != nil {
		_ = DB.Close()
	}
}
