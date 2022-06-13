package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"example.com/calorieCounter/food"
)

func Init() *gorm.DB {
    dbURL := "postgres://postgres:password@localhost:5432/postgres"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&food.Food{})

    return db
}