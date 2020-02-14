package main

import (
	"fmt"
	"math"
)

// interface

type Geometry interface {
	Name()
	Area() float32
	Perimeter() float32
}

// Rectangle

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Name() {
	fmt.Printf("{Rectangle(length=%f, breadth=%f}\n", r.Length, r.Breadth)
}

func (r Rectangle) Area() float32 {
	return (r.Length * r.Breadth)
}

func (r Rectangle) Perimeter() float32 {
	return (2 * (r.Length + r.Breadth))
}

// Circle

type Circle struct {
	Radius float32
}

func (c Circle) Name() {
	fmt.Printf("{Circle(radius=%f)}\n", c.Radius)
}

func (c Circle) Area() float32 {
	return float32(math.Pi * math.Pow(float64(c.Radius), 2))
}

func (c Circle) Perimeter() float32 {
	return (2 * math.Pi * c.Radius)
}
