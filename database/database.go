package database

import (
	"fmt"
	"log"
	"os"
	"sanjibook/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
		os.Exit(2)
	}

	log.Printf("Database has been connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Printf("Running migrations")
	db.AutoMigrate(&models.Tip{}, &models.Instruction{}, &models.Equipment{}, &models.Ingredient{}, &models.Recipe{}, &models.RecipesIngredients{})
	DB = Dbinstance{Db: db}
}
