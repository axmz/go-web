package main

import (
	"fmt"
)

func main() {
	// c := controller.NewController()
	// i := c.ShowController()
	// fmt.Println(i)
	// fmt.Printf("%p\n", &c)

	// c = controller.NewController()
	// i = c.ShowController()
	// fmt.Println(i)
	// fmt.Printf("%p\n", &c)

	// c = controller.NewController()
	// i = c.ShowController()
	// fmt.Println(i)
	// fmt.Printf("%p\n", &c)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	// x[1] = 999
	fmt.Println(len(x), cap(x))
	fmt.Printf("%p\n", x)

	y := x[1:3]
	fmt.Println(y)
	fmt.Println(len(y), cap(y))
	fmt.Printf("%p\n", y)

	x = append(x, 6)
	x[1] = 999
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	fmt.Printf("%p\n", x)
	fmt.Println(y)
	fmt.Println(len(y), cap(y))
	fmt.Printf("%p\n", y)
}
