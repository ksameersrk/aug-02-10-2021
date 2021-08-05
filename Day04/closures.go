package main

import "fmt"

//Closures
func main() {
	inc := incrementor()
	fmt.Printf("%T\n", inc)
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	inc2 := incrementor()
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())
}
func incrementor() func() int{
	x := 0
	return func() int {
		x++
		return x
	}
}