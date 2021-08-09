package main

import "fmt"

func main()  {
	//basics()
	//talkToDatabase()
}

func talkToDatabase() {
	fmt.Println("Open a connection to DB")
	defer fmt.Println("Close the connection to DB")
	fmt.Println("Execute a Query")
	fmt.Println("Done talking to Database")
}

func basics() {
	defer fmt.Println("First line in main")
	defer fmt.Println("Second line in main")
	defer func() {
		fmt.Println("Inside deferred function")
	}()
	fmt.Println("End of main")
	//The deferred stack is run
}
