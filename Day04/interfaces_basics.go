package main

import (
	"fmt"
	"strconv"
)
//class Truck implements Vehicle { testDrive(...) }
type Truck struct {
	model string
}

//class ElectricCar implements Vehicle { testDrive(...) }
type ElectricCar struct {
	model string
	power int
}

type Vehicle interface {
	testDrive()
	applyBrakes()
}

func main()  {
	tata := Truck { model: "Tata"}
	tesla := ElectricCar{model: "Tesla XYZ", power: 4000}
	//tata.testDriveTrucks()
	//tesla.testDriveElectricCar()
	testDriveVehicles(tata)
	testDriveVehicles(tesla)
}

func testDriveVehicles(v Vehicle) {
	v.testDrive()
	v.applyBrakes()
}

func (truck Truck) applyBrakes() {
	fmt.Println("Testing braking mechanism of " + truck.model)
}
func (car ElectricCar) applyBrakes() {
	fmt.Println("Testing braking mechanism of " + car.model)
}

//Truck implements Vehicle interface by having a receiver function for testDrive
func (truck Truck) testDrive()  {
	fmt.Println("On a test drive of the truck ", truck.model)
}

//ElectricCar implements Vehicle interface by having a receiver function for testDrive
func (car ElectricCar) testDrive()  {
	fmt.Println("On a test drive of EV ", car.model + " with power " + strconv.Itoa(car.power))
}
