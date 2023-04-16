package main

import "fmt"

func main() {
	arr1 := [3]int{7, 9, 4}
	slice1 := []int{7, 9, 4, 8}
	//cities := []string {"London", "NYC", "Colombo", "Tokyo"}

	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", slice1)

	slice1 = append(slice1, 88)

}
