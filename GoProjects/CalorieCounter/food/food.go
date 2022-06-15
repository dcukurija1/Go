package food
import "gorm.io/gorm"
type Food struct {
	gorm.Model
	Name string `json:"name"`
	Category string `json:"category"`
	Calories float32 `json:"calories"`
	PortionSize float32 `json:"portionSize"`
}

