package main

import (
	bl "Day05-Monday/bangalore"
	"Day05-Monday/intuit"
	"fmt"
	my "github.com/prabhu-sunderaraman/my-golang-libs"
)

func main()  {
	fmt.Println(my.GenerateRandomNumber())
	fmt.Println(intuit.PI)
	fmt.Println(intuit.GenerateRandomNumber())
	p1 := bl.Person{Name: "Joe", Age: 12}
	fmt.Println(p1)
}
