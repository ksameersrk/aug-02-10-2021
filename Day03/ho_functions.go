package main

import "fmt"

type EmptyFunction func()
type MathOperation func(int, int) int
type ComplexOperation func(int, int, int) (int, float32, float64)

func main()  {
	var work = func() {
		fmt.Println("Working")
	}
	var eat EmptyFunction = func() {
		fmt.Println("Eating")
	}
	doSomething(work)
	doSomething(eat)

	var add = func(a int, b int) int {
		return a + b
	}
	var multiply MathOperation = func(a int, b int) int {
		return a * b
	}
	math(add)
	math(multiply)
	randomGenerator()()
}

//func randomGenerator() func() {
func randomGenerator() EmptyFunction {
	fn := func() {
		fmt.Println(324.2343)
	}
	return fn
}

//math is a HO function
//func math(operation func(int, int) int) {
func math(operation MathOperation) {
	num1 := 10
	num2 := 20
	result := operation(num1, num2)
	fmt.Println(result)
}
//doSomething is a HO function
func doSomething(arg EmptyFunction)  {
//func doSomething(arg func())  {
	arg()
}
func doSomethingComplex(co ComplexOperation)  {

}
