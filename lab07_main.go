package main

// Define the following packages

import (
	"fmt"
	"./cars"
	"./common"
)

func main() {
	garage := make([]common.Vehicle, 3)

	garage[0], _ = cars.NewFord(2001)
	garage[1], _ = cars.NewHonda(2010)
	garage[2], _ = cars.NewBMW(2019)

	for _, v := range garage {
		//calls String() string; This is similar to our toString() method
		fmt.Println(v) //Prints model, year, type, seatingcapacity
	}
}
