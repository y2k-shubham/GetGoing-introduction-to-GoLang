package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Printf(sep string) {
	fmt.Printf("{Name: %s%sAge: %d%sModelNo: %d}\n", c.Name, sep, c.Age, sep, c.ModelNo)
}

func (c Car) Drive() {
	fmt.Println("driving")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name:    "chevy",
		Age:     8,
		ModelNo: 877,
	}
	fmt.Println(c)

	c.Printf("\t")
	c.Drive()
	fmt.Println(c.GetName())
}
