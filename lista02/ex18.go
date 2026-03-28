package main

import f "fmt"

func main() {
	var valor, tipo float64
	f.Println("Qual o valor normal do DVD?")
	f.Scan(&valor)
	f.Println("Qual a categoria?")
	f.Println("1 - Comum")
	f.Println("2 - Lançamento")
	f.Scan(&tipo)

	if tipo == 2 {
		valor += (valor * 15 / 100)
	}

	f.Printf("O preço do DVD é: R$ %.2f.", valor)
}
