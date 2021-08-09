package cars

import "errors"

type Car struct {
	model           string
	year            uint16
	seatingCapacity int
	carType         string
}

func (car Car) Model() string {
	return car.model
}

func (car Car) Year() uint16 {
	return car.year
}

func (car Car) SeatingCapacity() int {
	return car.seatingCapacity
}
func (car Car) Type() string {
	return car.carType
}

//WILL FIX THIS CONVERSION AFTER THE CLASS
func (car Car) String() string {
	return "Car: "+ car.model + ", (" + string(car.year) + "), seating capacity: " + string(car.seatingCapacity) + ", " + car.carType;
}

func NewHonda(year uint16) (Car, error) {
	var err error = nil
	if year < 2000 {
		err = errors.New("Year cannot be lt 2000")
	}
	return Car {model: "Honda Jazz", year: year, seatingCapacity: 5, carType: "Hatchback"}, err
}

func NewFord(year uint16) (Car, error) {
	return Car {model: "Ecosport", year: year, seatingCapacity: 6, carType: "SUV"}, nil
}

func NewBMW(year uint16) (Car, error) {
	return Car {model: "BMW", year: year, seatingCapacity: 5, carType: "Sedan"}, nil
}