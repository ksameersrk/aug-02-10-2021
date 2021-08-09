package main

import "fmt"

func main()  {
	defer catch()
	calculate()
	fmt.Println("******End of main")
}

func calculate() {
	//defer catch()
	//defer func() {
	//	rec := recover()
	//	if rec != nil {
	//		fmt.Println("Error", rec)
	//	}
	//} ()
	num := 10
	den := 0
	if den == 0 {
		panic("Don't divide by 0") //throw keyword
	}
	result := num / den
	arr := make([]int, 2)
	fmt.Println(arr[10])
	fmt.Println(result)
}

func catch() {
	if rec := recover(); rec != nil {
		fmt.Println("Error", rec)
	}
}