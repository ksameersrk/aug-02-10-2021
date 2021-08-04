package main

import "fmt"

func main()  {
	var greet = func() {
		fmt.Println("Hi")
	}
	greet()
	var add = func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(12, 34))
}

//func greet()  {
//	fmt.Println("Hi")
//}
