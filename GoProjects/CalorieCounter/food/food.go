package food
type Food struct {
	name string `json:"name"`
	category string `json:"category"`
	calories int `json:"calories"`
}
// setters
func (f *Food) SetName(name string){
	f.name = name
}

func (f *Food) SetCategory(c string){
	f.category = c
}

func (f *Food) SetCalories(c int){
	f.calories = c
}
// getters
func (f *Food) GetName() string {
	return f.name
}

func (f *Food) GetCategory() string {
	return f.category
}

func (f *Food) GetCalories() int {
	return f.calories
}

// builder pattern
// probably won't need it
func NewFood(name string, cat string, cal int) *Food{
	return &Food{name, cat, cal}
}

