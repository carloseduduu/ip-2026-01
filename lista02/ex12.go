package main

import f "fmt"

func main() {
	var idade int
	f.Print("Insira sua idade: ")
	f.Scan(&idade)
	if idade >= 0 && idade <= 2 {
		f.Println("Classificação: Recém-nascido")
	} else if idade >= 3 && idade <= 11 {
		f.Println("Classificação: Criança")
	} else if idade >= 12 && idade <= 19 {
		f.Println("Classificação: Adolescente")
	} else if idade >= 20 && idade <= 55 {
		f.Println("Classificação: Adulto")
	} else {
		f.Println("Classificação: Idoso")
	}
}
