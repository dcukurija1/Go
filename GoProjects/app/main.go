package main

import "fmt"

// example with languages
type greetings interface {
	hello()
}

type English struct {
	greet string
}
type German struct {
	greet string
}
type Bosnian struct {
	greet string
}

func (e English) hello() {
	fmt.Println("Hello on English: " + e.greet)
}
func (e German) hello() {
	fmt.Println("Hello on German: " + e.greet)
}
func (e Bosnian) hello() {
	fmt.Println("Hello on Bosnian: " + e.greet)
}

func sayHi(g greetings) {
	g.hello()
}

// a simple application
func main() {

	var eng English
	eng.greet = "Hello"
	var ger German
	ger.greet = "Hallo"
	var bos Bosnian
	bos.greet = "Pozdrav"

	sayHi(eng)
	sayHi(ger)
	sayHi(bos)
}
