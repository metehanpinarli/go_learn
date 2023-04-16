package main

import "fmt"

func main() {
	//fmt.Println("Hello World test 2")
	//welcome("50")
	a := "bir"
	{
		a := "iki"
		fmt.Println(&a)
	}
	fmt.Println(&a)
}

func welcome(name string) string {
	result := "Hoşgeldin" + name
	return result
}
