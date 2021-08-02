package main

import (
	"fmt"
)

func main()  {
	count := 0
	//while loop in languages like Java, C#, JS
	for count < 5 {
		fmt.Println(count)
		count++
	}
	//for {
	//	fmt.Println("infinite loop")
	//}

	//Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}

	for x := 0; x < 5; {
		fmt.Println(x)
		if x == 3 {
			break
		}
		x++
	}

}