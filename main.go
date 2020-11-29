package main

import (
	"fmt"
)

type Custom struct {
	name string
}

func main() {
	c := Custom{name: "custard"}
	c.change()
	fmt.Println(c.name)
}

func (c Custom) change() {
	c.name = "changed"
}
