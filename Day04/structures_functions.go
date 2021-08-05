package main

import "fmt"

type Book struct {
	title string
	price float32
	read //NOT COMMONLY USED
}
type read func()
func main()  {
	b1 := Book{title: "Ajax hacks",price: 34.45, read: func() { fmt.Println("Reading")} }
	//printDetails(b1)
	b1.printDetails()
	b1.read() //Looks like a receiver function, but it's a member of the struct//NOT VERY COMMONLY used

	b1.buy("Amazon")

	library := []Book {
		Book{title: "A", price: 12.34},
		Book{title: "B", price: 45.67},
	}
	for i := 0; i < len(library); i++ {
		book := library[i]
		fmt.Println(book)
	}

	emptyBook := Book {title: "ABC"}
	fmt.Println(emptyBook)
	initializePrice(emptyBook)
	fmt.Println(emptyBook)
}
//By default structs are also passed by value ONLY
func initializePrice(b Book)  {
	b.price = 45.56
}

func (b Book) buy(store string)  {
	fmt.Println("Buying " + b.title + " from " + store)
}

func (b Book) printDetails() {
	fmt.Println(b.title, b.price)
}
