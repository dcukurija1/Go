package food
import "fmt"
type Plate struct {
	FoodItems []Food `json:"foodItems"`
}

func (p *Plate)AddToPlate(f Food) {
	p.FoodItems = append(p.FoodItems, f)
}

func (p Plate)CurrentCalories() float32 {
	var sum float32 = 0
	for _, food := range p.FoodItems {
		sum += food.Calories
	} 
	return sum
}

func (p Plate) ShowPlate() {
	fmt.Println("Food on your plate: ")
	for _, food := range p.FoodItems {
		fmt.Println(food.Name, food.Calories, "calories in", food.PortionSize,"g")
	}
	fmt.Println()
}

func (p Plate) Find(f Food) bool {
	for _, food := range p.FoodItems {
		if f == food {
			return true
		}
	}
	return false
}