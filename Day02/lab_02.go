package main

import (
	"fmt"
	"strconv"
	"strings"
)

var exclusionList = []int {10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910}
var input = "+5 -1 +9 +5 -67 +5 -3 +2 -4 +6 +8 -13 +2 -4 +6 +18 -3 +2 -4 +6 +88 +15 -1 +9 +5 -67 +45 -3 +2 -4 +36 +8 -13 +2 -4 +6 +18 -3 +2 -74 +11 +109"

func main()  {
	calculateSumOfOddNumbers()
	calculateIntegerAverage()
	computeSumOfInput()
}

func computeSumOfInput() {
	items := strings.Split(input, " ")
	sum := 0
	for _, value := range items {
		num, err := strconv.Atoi(value)
		if err == nil {
			sum += num
		}
	}
	fmt.Println("Sum", sum)
}

func calculateIntegerAverage() {
	sum := 0
	for i := 0; i <= 10001; i++ {
		if !numberNotInExclusionList(i) {
			sum += i
		}
	}
	average := sum / (10001 - (len(exclusionList))) //BRACKETS work in Golang!!!
	fmt.Println("Average", average)
}

func numberNotInExclusionList(num int) bool {
	found := false
	for _, value := range exclusionList {
		if value == num {
			found = true
			break
		}

	}
	return found
}



func calculateSumOfOddNumbers() {
	sum := 0
	for i := 1; i < 10001; i += 2 {
		sum += i
	}
	fmt.Println("Sum of odd numbers", sum)
}
