package main

import (
	"fmt"
	"strings"
)

type currency float32
type name string
func main()  {

	//var dollar float32 = 73.45
	//var euro float32 = 89.34
	var dollar currency = 73.45
	var euro currency = 89.34
	var pi float32 = 3.14
	fmt.Println(dollar, euro, pi)

	//OO flavour
	//printCurrency(dollar)
	//printCurrency(euro)
	dollar.printCurrency()
	euro.printCurrency()
	dollar.value()
	euro.value()
	var goldMedalist name = "Joe"
	var silverMedalist name = "Sindhu"

	goldMedalist.toU()
	silverMedalist.toU()
}
func (s name) toU()  {
	fmt.Println(strings.ToUpper(string(s)))
}
func (c currency) value()  {
	fmt.Println(c)
}

//receiver function
func (c currency) printCurrency()  {
	fmt.Println(c)
}
