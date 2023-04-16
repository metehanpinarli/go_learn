package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = create("deneme/dosya/deneme.txt")
	//fmt.Println(newFile)
	file := getInfo("deneme/dosya")
	fmt.Println(file.Name())
	fmt.Println(file.Mode())
	fmt.Println(file.IsDir())
	fmt.Println(file.ModTime())
	fmt.Println(file.Sys())
	fmt.Println(file.Sys())

}

func create(p string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(p), 0770)
	if err != nil {
		return nil, err
	}
	return os.Create(p)
}

func getInfo(p string) os.FileInfo {
	info, err := os.Stat(p)
	if err != nil {
		log.Fatal("hata")
	}
	return info

}
