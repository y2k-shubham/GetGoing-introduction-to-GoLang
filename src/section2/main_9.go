package main

import "fmt"

func continue_demo() {
	fmt.Println("continue demo")
	for i := 1; i <= 10; i++ {
		if ((i % 2) == 0) {
			continue
		}
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
}

func break_demo() {
	fmt.Println("break demo")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
}

func switch_case_demo_1(day string) {
	switch day {
	case "friday", "saturday": fmt.Printf("%s => yay!\n", day)
	case "sunday": fmt.Printf("%s => oh no; its gone!\n", day)
	default:
		fmt.Printf("%s => meh!\n", day)
	}
}

func switch_case_demo_2(day string) {
	switch {
	case day == "friday" || day == "saturday": fmt.Printf("%s => yay!\n", day)
	case day == "sunday": fmt.Printf("%s => oh no; its gone!\n", day)
	default:
		fmt.Printf("%s => meh!\n", day)
	}
}

func main() {
	continue_demo()
	break_demo()

	fmt.Println("switch_case_demo_1")
	switch_case_demo_1("monday")
	switch_case_demo_1("tuesday")
	switch_case_demo_1("wednesday")

	fmt.Println("switch_case_demo_2")
	switch_case_demo_2("thursday")
	switch_case_demo_2("friday")
	switch_case_demo_2("saturday")
	switch_case_demo_2("sunday")
}
