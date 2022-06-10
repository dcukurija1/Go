package food
import "fmt"
type Plate struct {
	plate []Food
}
type plate_int interface{
	AddToPlate()
	CurrentCalories()
	ShowPlate()
} 
func (p Plate)AddToPlate(f Food) {
	p.plate = append(p.plate, f)
}

func (p Plate)CurrentCalories() int {
	sum := 0
	for _, food := range p.plate {
		sum += food.Calories
	} 
	return sum
}

func (p Plate) ShowPlate() {
	for _, food := range p.plate {
		fmt.Println(food.Name, " ", food.Category, " ", food.Calories)
	} 
}