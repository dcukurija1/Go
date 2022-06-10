package food
import "fmt"
type Plate struct {
	Items []Food
}
type plate_int interface{
	AddToPlate()
	CurrentCalories() int
	ShowPlate()
} 
func (p Plate)AddToPlate(f Food) {
	p.Items = append(p.Items, f)
}

func (p Plate)CurrentCalories() int {
	sum := 0
	for _, food := range p.Items {
		sum += food.Calories
	} 
	return sum
}

func (p Plate) ShowPlate() {
	for _, food := range p.Items {
		fmt.Println(food)
	} 
}