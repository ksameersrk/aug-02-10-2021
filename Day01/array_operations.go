package main

import "fmt"

func main() {
	//Slices
	//shallowCopy()
	//deepCopy()
	var numbers = []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	var myList = make([]int, 3)
	fmt.Println(myList, numbers)
	//var lst = append(myList, numbers...) //spread operator in JavaScript
	var lst = append(myList, numbers[:2]...)
	fmt.Println(lst)
	fmt.Printf("%p %p %p\n", &lst, &myList, &numbers)
}

func deepCopy() {
	var numbers = []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	//var first3Numbers = numbers[0:3]
	var first3Numbers = make([]int, 3)
	//var first3Numbers [3]int //ERROR - [3]int <> []int
	//var first3Numbers = []int{0, 0, 0}
	copy(first3Numbers, numbers[0:3]) //deep copy
	first3Numbers[0] = 100000
	fmt.Println(first3Numbers, numbers)
}

func shallowCopy() {
	var numbers = []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	var first3Numbers = numbers[0:3]
	first3Numbers[0] = 1000 //Modifying this modifies the original source as well
	fmt.Println(numbers, first3Numbers)

	var last3Numbers = numbers[6:]
	fmt.Println(last3Numbers)

	var myNumbers = numbers //shallow
	myNumbers[0] = 100000
	fmt.Println(myNumbers, numbers)

	newNumbersList := numbers //shallow
	newNumbersList[0] = 999999
	fmt.Println(newNumbersList, numbers)
}
