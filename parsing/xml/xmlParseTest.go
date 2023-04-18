package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Members struct {
	//XMLName    xml.Name `xml:"members"`
	Members []Member `xml:"member"`
}
type Member struct {
	//XMLName   xml.Name `xml:"member"`
	Type   string `xml:"type,attr"`
	Name   string `xml:"name"`
	Social Social `xml:"social"`
}
type Social struct {
	//XMLName     xml.Name `xml:"social"`
	Facebook string `xml:"facebook"`
	Twitter  string `xml:"twitter"`
	Youtube  string `xml:"youtube"`
}

func main() {

	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	//Yerleştirme işlemi için değişken oluşturuyoruz.
	var members Members
	xml.Unmarshal(byteValue, &members)
	fmt.Println(members.Members[0].Social)
}
