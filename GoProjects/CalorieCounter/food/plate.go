package food
import "fmt"
type Plate struct {
	FoodItems []Food `json:"fooditems"`
}

func (p *Plate)AddToPlate(f Food) {
	p.FoodItems = append(p.FoodItems, f)
}

func (p Plate)CurrentCalories() int {
	sum := 0
	for _, food := range p.FoodItems {
		sum += food.GetCalories()
	} 
	return sum
}

func (p Plate) ShowPlate() {
	fmt.Println("Food on your plate: ")
	for _, food := range p.FoodItems {
		fmt.Println(food.GetName(), food.GetCalories(), "calories")
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