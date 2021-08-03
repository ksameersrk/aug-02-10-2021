package main

import "fmt"

func main()  {
	//[size]datatype for arrays
	//map[datatype]datatype

	var romanNumbers map[int]string = make(map[int]string)
	romanNumbers[1] = "I"
	romanNumbers[2] = "II"
	romanNumbers[3] = "III"
	delete(romanNumbers, 2)

	for key, value := range romanNumbers {
		fmt.Println(key, value)
	}
	fmt.Println(romanNumbers, len(romanNumbers))

}