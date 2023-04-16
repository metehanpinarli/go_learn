package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var number = 50
	println(number, reflect.TypeOf(number))
	println(reflect.TypeOf(number))
	str := strconv.Itoa(number)
	println(str, reflect.TypeOf(str))

	var (
		height = 20
		with   = 80
	)
	fmt.Println(height, with)

	var isCurrent bool

	isCurrent = true

	println(isCurrent)

	text := "mete"

	println(text)

}
