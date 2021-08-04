package main

import "fmt"

var numbers = []int {101, 22, 43, 14, 5, 906, 310, 561, 84, 23, 100}
type Finder func(int) bool
type Computation func(int) int

func main()  {
	partI()
	partII()
}

func partII() {
	var sumOfSquares Computation = func(num int) int {
		return num * num
	}
	fmt.Println(sumOfSquares.compute())

	var sumOfCubes Computation = func(num int) int {
		return num * num * num
	}
	fmt.Println(sumOfCubes.compute())
}

func (logic Computation) compute() (result int)  {
	for i := 0; i < len(numbers); i++ {
		result += logic(numbers[i])
	}
	return
}

func partI() {
	even := func(num int) bool {
		return num%2 == 0
	}
	lstEven := find(even)
	fmt.Println(lstEven)

	lstOdd := find(func(num int) bool {
		return num%2 != 0
	})
	fmt.Println(lstOdd)

	lstDivBy5 := find(func(num int) bool {
		return num%5 == 0
	})
	fmt.Println(lstDivBy5)
}

func find(predicate Finder) (result []int) {
	result = make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		if predicate(num) {
			result = append(result, num)
		}
	}
	return
}