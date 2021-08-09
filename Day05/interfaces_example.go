package main

import "fmt"

type Circle struct {
	radius float32
}
type Square struct {
	side float32
}
type Shape interface {
	area() float32
}

func main()  {
	var cir Shape = Circle{radius: 23.45}
	var squ Shape = Square{side: 12.34}
	var shapes []Shape = []Shape {cir, squ}

	for i := 0; i < len(shapes); i++ {
		fmt.Println(shapes[i].area())
	}
}

func (square Square) area() float32  {
	return square.side * square.side
}

func (circle Circle) area() float32 {
	return 3.14 * circle.radius * circle.radius
}