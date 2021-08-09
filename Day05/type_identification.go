package main

import "fmt"

type Rect struct {

}
func main()  {
	//x := 65536 * 65536 * 65536 * 65536
	//fmt.Printf("%T\n\n", x)
	//doSomething(x)
	doSomething("monday")
	doSomething(false)
	doSomething(Rect{

	})
}

//Equivalent to RTTI(instanceof) in Java
func doSomething(any interface{})  {
	switch any.(type) {
	case Rect:
		var data Rect = any.(Rect)
		fmt.Println("It's a rectangle", data)
	case int:
		var data int = any.(int)
		data++
		fmt.Println("It's an integer", data)
	case string:
		var data string = any.(string)
		fmt.Println("It's a string", data)
	default:
		fmt.Printf("%T\n", any)
	}
}
