package main

import "fmt"

type Salca interface {
	Ye()
	Yesme()
}

type biberSalcasi struct {
	//buraya birşey yazmamız gerekli değil
}

func (s biberSalcasi) Yesme() {
	//TODO implement me
	panic("implement me")
}

func (s biberSalcasi) Ye() {
	fmt.Println("Biber salçası yenildi")
}

func main() {
	var salcaInterface Salca
	biber := biberSalcasi{}

	salcaInterface = &biber
	salcaInterface.Yesme()
}
