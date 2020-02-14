package main

import (
	"errors"
	"fmt"
)

// factory

func CreateGeometry(name string) (Geometry, error) {
	switch name {
	case "rectangle":
		return Rectangle{Length: 4, Breadth: 3}, nil
	case "circle":
		return Circle{Radius: 1.5}, nil
	default:
		return nil, errors.New(fmt.Sprint("Invalid name '%s'", name))
	}
}


func main() {
	var r, _ = CreateGeometry("rectangle")
	r.Name()
	fmt.Printf("Rectangle perimeter = %f\n", r.Perimeter())
	fmt.Printf("Rectangle area      = %f\n", r.Area())

	// circle
	var c, _ = CreateGeometry("circle")
	c.Name()
	fmt.Printf("Circle perimeter = %f\n", c.Perimeter())
	fmt.Printf("Circle area      = %f\n", c.Area())
}
