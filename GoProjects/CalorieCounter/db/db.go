package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"example.com/calorieCounter/food"
)

func initialFoodItems() []food.Food {
	return []food.Food{*food.NewFood(1, "Apple", "Fruit", 50, 100),
					   *food.NewFood(2, "Orange", "Fruit", 60, 100),
					   *food.NewFood(3, "Pear", "Fruit", 65, 100)}
}

func Init() *gorm.DB {
    dbURL := "postgres://postgres:password@localhost:5432/postgres"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&food.Food{})
	/*var foods = initialFoodItems()

	db.Create(&foods)
	// testing db
	for _,f := range foods {
		log.Println(f)
	}*/

    return db
}