package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type currency float32
type name string
type age int32
type coll []int

func main()  {
	var i age = 12
	strconv.Itoa(int(i))
	var lst coll = []int {1}
	fmt.Printf("%T\n", lst)
	sort.SearchInts(lst, int(i))

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
