package main

import "fmt"

func main() {
	var result int = addition(5, 7, 8, 10)
	fmt.Println(result)
}
func addition(numbers ...int) int {
	var result int
	for _, n := range numbers {
		result += n
	}
	return result
}
