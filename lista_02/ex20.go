package main

import f "fmt"

func main() {
	var valor float64
	var tipo int
	f.Println("Qual o valor do produto?")
	f.Scan(&valor)

	f.Println("Qual a forma de pagamento?")
	f.Println("1 - À vista, dinheiro ou cheque, 10% de desconto")
	f.Println("2 - À vista, cartão de crédito, 5% de desconto")
	f.Println("3 - Em 2 vezes, preço normal de etiqueta sem juros")
	f.Println("4 - Em 3 vezes, preço normal de etiqueta + 10% de juros")
	f.Scan(&tipo)

	switch tipo {
	case 1:
		valor += -(valor * 10 / 100)
	case 2:
		valor += -(valor * 5 / 100)
	case 4:
		valor += valor * 10 / 100
	}

	f.Printf("O valor do produto é: R$ %.2f", valor)
}
