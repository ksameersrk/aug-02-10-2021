package main

import "fmt"

//pi := 3.14
var pi float32 = 3.14
func variables()  {
	var i int = 10
	var j = 20
	fmt.Println(i, j)

	//var name string = "Sam"
	//var name = "Sam"
	name := "Sam"
	fmt.Println(name)
	fmt.Printf("%T\n", name)

	inStock := true
	letter := 'a' //int32
	fmt.Printf("%T, %T\n", inStock, letter)
}
