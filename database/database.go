package database

import (
	"log"
	"os"

	"github.com/SourabhG16/goProject1/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("Api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Faile to connetc to database \n Error is :", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to datbase successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}
}
