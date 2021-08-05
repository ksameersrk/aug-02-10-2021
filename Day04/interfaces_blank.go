package main

import "fmt"

type ANY interface {

}
type Cat struct {}
func main()  {
	var data1 ANY = 120
	var data2 ANY = true
	fmt.Println(data1, data2)

	var cat Cat = Cat{}
	doSomething(cat)

	var x int = 10
	doSomething(x)

	doSomething("hello, what is going on?")
	printToConsole(1)
	printToConsole(true)
	printToConsole(1, 23.4, true, "hello", []int{})

	var lst []interface{} = []interface{} {1, 12, true, "string"}
	fmt.Println(lst)
}

func printToConsole(e ... interface{})  {
	fmt.Println(e)
}

func doSomething(e ANY)  {
	fmt.Printf("%T\n", e)
}
