package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Part-1")
	Anything(3.142)
	Anything("Pi")
	Anything(struct {
		Nr int
		Dr int
	}{Nr: 22, Dr: 7})

	fmt.Println("Part-2")
	var myMap map[string]interface{} = make(map[string]interface{})
	myMap["my-puchi-1"] = 123
	myMap["123"] = "my-puchi-1"
	myMap["my-puchi-2"] = struct {
		Puchi string
		Kuchi int
	}{Puchi: "my-puchi-2", Kuchi: 22}
	fmt.Println(myMap)
}
