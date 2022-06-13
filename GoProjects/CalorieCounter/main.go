package main

import (
	"fmt"
	"example.com/calorieCounter/food"
)

func main() {
	var p food.Plate
	p.AddToPlate(*food.NewFood("Apple", "Fruit", 50));
	p.ShowPlate()
	p.AddToPlate(*food.NewFood("Orange", "Fruit", 60));
	p.ShowPlate()
	fmt.Println("Total calories: ", p.CurrentCalories())
}
