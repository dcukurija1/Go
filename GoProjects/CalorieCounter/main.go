package main

import (
	"fmt"
	"example.com/calorieCounter/food"
)

func main() {
	var p food.Plate
	p.AddToPlate(food.Food{"Apple", "Fruit", 50,});
	p.ShowPlate()
	p.AddToPlate(food.Food{"Orange", "Fruit", 60,});
	p.ShowPlate()
	fmt.Println(p.CurrentCalories())
}
