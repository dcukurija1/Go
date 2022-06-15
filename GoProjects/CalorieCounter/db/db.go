package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"example.com/calorieCounter/food"
)

func initialFoodItems() []food.Food {
	return []food.Food{food.Food{ Name:"Apple", Category:"Fruit", Calories: 50, PortionSize: 100},
					   food.Food{ Name:"Green Apple", Category:"Fruit", Calories: 50, PortionSize: 100},
					food.Food{ Name:"Red Apple", Category:"Fruit", Calories: 50, PortionSize: 100}}
}

func Init() *gorm.DB {
    dbURL := "postgres://postgres:password@localhost:5432/postgres"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }
    
    if !db.Migrator().HasTable(&food.Food{}) {
        db.AutoMigrate(&food.Food{})
	    db.Create(initialFoodItems())
    }
    return db
}