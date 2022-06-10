package hello

import (
	"fmt"
)

func Hi() string {
	defer func() {
		fmt.Println("napisi nesto")
	}()
	fmt.Println("poslije defer")
	return "Hello world"
}

func Div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("dividing with 0 not possible")
	}
	return a / b, nil
}
