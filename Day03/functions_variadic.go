package main

func main()  {
	sum(1, 2, 3)
	myList := []int {120, 20, 324, 4354}
	sum(myList ...)//Spread out the array and pass the numbers
	addition(myList)
}

func addition(numbers []int) int {
	return 0
}

func sum(numbers ... int) int  {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}