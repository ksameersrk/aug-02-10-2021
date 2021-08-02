package main

import "fmt"

func main()  {
	num := 10

	//if-else paranthesis is not required if (num % 2 == 0)
	//Positioning of curly braces and else keyword is strict
	if num % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if i := 5; i > 3 {
		fmt.Println("i is gt 3")
	}
	//fmt.Println(i) //Not visible


	switch num % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("I don't know what number this is")
	}

	switch x := 14; x > 10 {
	case true:
		fmt.Println("x is gt 10 I")
		fmt.Println("x is gt 10 II")
	case false:
		fmt.Println("x is lt 10")
	}

	switch day := "Tuesday"; day {
	case "Tuesday":
		fmt.Println("Huh!!!")
	case "Wednesday", "Thursday":
		fmt.Println("Long week")
	case "Friday":
		fmt.Println("TGIF")
		//fallthrough
	case "Saturday", "Sunday":
		fmt.Println("whohohohoh")
	default:
		fmt.Println("Life sucks! It's just Monday")
	}

}
