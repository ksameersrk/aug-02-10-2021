package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	name := "Sam"
	fmt.Println(name, string(name[0]), string(name[1]))
	fmt.Println(len(name))
	//fmt.Println(name[4]) //Error

	fmt.Println(strings.ToUpper(name), strings.ToLower(name))

	city := "Paris Hilton"
	newString := strings.ReplaceAll(city, " ", "SPACE")
	fmt.Println(newString)

	letter := 'A'
	fmt.Println(string(letter))

	paragraph := `This 
		is a
			paragraph`
	fmt.Println(paragraph)

	country := "India"
	chars := strings.Split(country, "")
	for i := 0; i < len(chars); i++ {
		fmt.Println(string(chars[i]))
	}

	fmt.Printf("Type of chars is %T\n", chars)

	i := "12"
	num, _ := strconv.Atoi(i) //INTENTIONALLY NOT GOING TO EXPLAIN TODAY
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	//a, _, _ := myFunc()
}
