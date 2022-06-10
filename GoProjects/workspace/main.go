package main

import (
	"encoding/json"
	"fmt"

	"example.com/hello"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) ispis() {
	fmt.Println("Name: ", p.Name, "\nAge: ", p.Age)

}
func main() {
	p := Person{
		Name: "Konj",
		Age:  21,
	}
	//p.ispis()
	pjson, err := json.Marshal(p)
	fmt.Println(string(pjson), err)
	var p2 Person
	json.Unmarshal(pjson, &p2)
	p2.ispis()
	if err2 := json.Unmarshal(pjson, &p2); err2 == nil {
		fmt.Println(hello.Hi())
	}

}
