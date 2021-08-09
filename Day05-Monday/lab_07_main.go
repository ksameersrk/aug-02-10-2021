package main

// Define the following packages

import (
	"Day05-Monday/cars"
	"Day05-Monday/common"
	"fmt"
)

func main() {
	_, err := cars.NewHonda(1999)
	if err != nil {
		fmt.Println("Error:", err)
	}
	garage := make([]common.Vehicle, 3)

	garage[0], _ = cars.NewFord(2001)
	garage[1], _ = cars.NewHonda(2010)
	garage[2], _ = cars.NewBMW(2019)

	for _, v := range garage {
		//calls String() string; This is similar to our toString() method
		fmt.Println(v) //Prints model, year, type, seatingcapacity
	}
}
