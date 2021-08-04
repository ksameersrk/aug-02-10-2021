package main

import "fmt"

func main()  {
	sq, cu := operate(4)
	fmt.Println(sq, cu)

	_, cube := operate(5)
	fmt.Println(cube)

	name, age, _ := getDetails()
	fmt.Println(name, age)
}

func getDetails() (name string, age int, salary float32) {
	name = "Sam"
	age = 12
	salary = 32487
	return
}

func operate(num int) (int, int) {
	square := num * num
	cube := num * num * num

	return square, cube
}


func square(num int) int  {
	return num * num
}