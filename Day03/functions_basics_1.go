package main

import "fmt"

func main()  {
	sum := add(12, 12)
	diff := subtract(12, 12)
	fmt.Println(sum, diff)
}

func add(a int, b int) int  {
	return a + b
}

//Named functions
func subtract(a int, b int) (result int) {
	result = a - b
	return //Mandatory
}

