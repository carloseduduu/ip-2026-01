package main

import f "fmt"

func main() {
	var n1 int
	f.Print("INFORME UM NÚMERO INTEIRO: ")
	f.Scan(&n1)
	if n1 > 20 && n1 < 90 {
		f.Println("O valor está entre 20 e 90")
	} else {
		f.Println("O valor não está entre 20 e 90")
	}
}
