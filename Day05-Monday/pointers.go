package main

import "fmt"

type Class struct {
	name string
}

func main()  {
	//var nameOfPointer *datatype

	x := 10
	var ptrToX *int = &x
	fmt.Println(*ptrToX)
	*ptrToX = 11
	fmt.Println(x)
	x++
	fmt.Println(*ptrToX)
	//ptrToX++ //Pointer arithmetic is NOT allowed in GO
	y := 11
	ptrToX = &y
	*ptrToX++
	fmt.Println(x)

	cls := Class{}
	initializeClass(&cls)
	fmt.Println(cls)

	arr := [10]int {}
	var arrPtr *[10]int = &arr
	fmt.Println(*arrPtr, arrPtr, ptrToX)
}
func initializeClass(cls *Class)  {
	cls.name = "Golang"
}