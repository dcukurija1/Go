package food
type Food struct {
	id int `json:"id" gorm:"primaryKey"`
	name string `json:"name"`
	category string `json:"category"`
	calories float32 `json:"calories"`
	portionSize float32 `json:"portionSize"`
}
// setters
func (f *Food) SetId(id int) {
	f.id = id
}
func (f *Food) SetName(name string){
	f.name = name
}

func (f *Food) SetCategory(category string){
	f.category = category
}

func (f *Food) SetCalories(calories float32){
	f.calories = calories
}

func (f *Food) SetPortionSize(size float32){
	f.portionSize = size
}
// getters
func (f *Food) GetId() int {
	return f.id
}

func (f *Food) GetName() string {
	return f.name
}

func (f *Food) GetCategory() string {
	return f.category
}

func (f *Food) GetCalories() float32 {
	return f.calories
}

func (f *Food) GetPortionSize() float32 {
	return f.portionSize
}

// builder pattern
func NewFood(id int, name string, cat string, cal float32, size float32) *Food{
	return &Food{id, name, cat, cal, size}
}

