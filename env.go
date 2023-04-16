package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Args {
		fmt.Println(env)
	}
}
