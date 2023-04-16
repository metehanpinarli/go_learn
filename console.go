package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var message string = "Please enter your age"
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	length := len(value)
	fmt.Println(strconv.Itoa(length))

}
