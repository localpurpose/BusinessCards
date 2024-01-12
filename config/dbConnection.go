package config

import (
	"fmt"
	"github.com/buscard/models"
	"github.com/gofiber/fiber/v2/log"
	env "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	//	///	///	///	///	///	///	///	///	///	///
	//	Database Connection
	//	Loading Env. Variables, gets it inside.
	//	///	///	///	///	///	///	///	///	///	///

	err := env.Load()

	dbHost := os.Getenv("HOST")
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("USER")
	dbPwd := os.Getenv("PASSWORD")

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection FAIL")
	}

	DB = db
	AutoMigrate(db)
	log.Info("DB Connection OK")
	log.Info("DB Migrations OK")
}

func AutoMigrate(conn *gorm.DB) {
	conn.Debug().AutoMigrate(
		&models.User{},
	)
}
