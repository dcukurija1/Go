package main

import (
	"fmt"
	"example.com/calorieCounter/food"
)

func main() {
	var p food.Plate
	apple := food.Food{"Apple", "Fruit", 50,}
	//fmt.Println(apple)
	orange := food.Food{"Orange", "Fruit", 60,}
	p.AddToPlate(apple);
	p.Items = append(p.Items, apple)
	p.ShowPlate()
	p.AddToPlate(orange);
	p.ShowPlate()
	fmt.Println(p.CurrentCalories())
}
