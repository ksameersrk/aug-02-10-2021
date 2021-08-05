package main

import "fmt"

type Person struct {
	name string
	age int
	isSingle bool
	myCar Car
	Brain //anonymous field; the properties can be accessed directly; Exported field
	//MathematicalBrain
}
type MathematicalBrain struct {
	iq int
}
type Brain struct {
	iq int
}
type Car struct {
	model string
	Engine
}
type Engine struct {
	power int
}
func main()  {
	var p1 Person
	fmt.Println(p1)

	var p2 Person = Person{name: "John", age: 12, isSingle: true, myCar: Car{model: "BMW"}}
	fmt.Println(p2.name, p2.age, p2.isSingle, p2.myCar.model)
	p3 := Person{
		name: "Ram",
		age: 20,
		myCar: Car{
			model: "Nano",
		},
	}
	p3.iq = 100 //iq is present in Brain type and it's exported
	p3.myCar.power = 2000 //power is present in Engine type and it's exported
	fmt.Println(p3)
}
