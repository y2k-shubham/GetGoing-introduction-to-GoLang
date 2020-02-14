package main

import (
	"fmt"
)


// main
func main() {
	// rectangle
	var r Geometry = Rectangle{Length: 4, Breadth: 3}
	r.Name()
	fmt.Printf("Rectangle perimeter = %f\n", r.Perimeter())
	fmt.Printf("Rectangle area      = %f\n", r.Area())

	// circle
	var c Geometry = Circle{Radius: 1.5}
	c.Name()
	fmt.Printf("Circle perimeter = %f\n", c.Perimeter())
	fmt.Printf("Circle area      = %f\n", c.Area())
}
