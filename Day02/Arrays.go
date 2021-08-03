package main

import (
	"fmt"
)

func main()  {
	basics()
	//for i := 0; i < (100 + 5)/6; i++ {
	//
	//}
}

func basics() {
	//[size]datatype
	var numbers [10]int
	numbers[0] = 100
	numbers[9] = 900
	fmt.Println(numbers, len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	//String[] names = new String[] {"Sam", "Ram", "Mary"}; in Java
	var names = []string {"Sam", "Ram", "Mary"}
	fmt.Println(names, len(names))

	var arr []int
	//arr[0] = 10 //ERROR
	if arr == nil {
		fmt.Println(len(arr))
	}
	numberOfShirts := 10
	var shirts = make([]string, numberOfShirts) //make is used initialize an array dynamically
	//delete()
	fmt.Println(shirts)

	var even = []int {2, 4, 6, 8, 10}

	//Similar to a for-each or for-in other languages
	for index, value := range even {
		fmt.Println(index, value)
	}
	for _, val := range even {
		fmt.Println(val)
	}
}
