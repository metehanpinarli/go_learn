package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = "1"
	var x = 5
	var y float32 = 5.6

	fmt.Println(myString, x, y)

	//String to Int
	number, _ := strconv.Atoi(myString)
	println(number)
	result := number + 2
	println(result)

	//int to String

	_string := strconv.Itoa(x)

	println(_string)

	//Casting

	var i int8 = -10
	var f float64 = float64(i)
	var u uint = uint(f)

	var test uint8 = uint8(i)

	fmt.Println(i, f, u, test)

}
