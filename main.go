package main

import (
	"fmt"
)

type Greeter struct {
	helloPhrase string
}

func (g Greeter) Hello() {
	fmt.Println(g.helloPhrase)
}

func main() {
	g := Greeter{helloPhrase: "Hey everyone!"}
	g.Hello()
}