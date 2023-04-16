package main

import "fmt"

type Car struct {
	color       string
	maxSpeed    int
	isAutomatic bool
}

func newCarGenerate() *Car {
	car := new(Car)
	return car
}

func (car Car) start() {
	fmt.Printf("%v started", car)
}

func main() {

	var test Car = Car{isAutomatic: true, maxSpeed: 200, color: "Purple"}
	test.isAutomatic = false

	myCar := Car{color: "Blue", maxSpeed: 190, isAutomatic: true}
	//myCarPointer := &myCar
	fmt.Printf("MY Car:%v", myCar)
	fmt.Println()
	fmt.Println(&myCar.color)

	newCar := newCarGenerate()
	newCar.maxSpeed = 120
	fmt.Printf("%v\n", newCar)

	newCar2 := newCarGenerate()
	newCar2.maxSpeed = 150
	fmt.Printf("%v\n", newCar2)
	newCar2.start()

}
