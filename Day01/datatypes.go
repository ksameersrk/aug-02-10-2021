package main

import "fmt"

var company string
func main()  {
	//int, bool, string, character, float
	//var nameOfTheVariable datatype

	//Variables have to USED
	var i int = 10
	var b bool = true
	var str string = "Hello"
	var pi float32 = 3.14

	//Omit the datatype as well. Type inference
	var j = 12

	fmt.Println(i, b, str, pi, j)

	var age int
	var name string
	var isMarried bool
	fmt.Println(age, name, isMarried)

	var ch = 'A'
	fmt.Println(string(ch)) //Typecast a char to string
}
