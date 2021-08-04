package main

import (
	"fmt"
	"sort"
)

type MyNumbers []int

func main()  {
	var myCollection MyNumbers = []int {10, 200, 34, 78, 9}
	fmt.Println(myCollection.sum())
	fmt.Println(myCollection.min())
	fmt.Println(myCollection.max())
}

func (lst MyNumbers) sum() (total int) {
	for _, number := range lst {
		total += number
	}
	return
}

func (lst MyNumbers) max() int {
	sort.Ints(lst)
	return lst[len(lst) - 1]
}

func (lst MyNumbers) min() int {
	sort.Ints(lst)
	return lst[0]
}