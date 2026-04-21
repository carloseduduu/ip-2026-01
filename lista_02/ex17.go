package main

import f "fmt"

func main() {
	var valor, metros float64
	var conta, tipo int
	f.Print("Insira o número da conta: ")
	f.Scan(&conta)
	f.Print("Insira seu consumo: ")
	f.Scan(&metros)
	f.Println("Insira sua categoria:")
	f.Println("1 - Residencial;")
	f.Println("2 - Comercial;")
	f.Println("3 - Industrial.")
	f.Scan(&tipo)

	// Loop caso o imput não for válido
	for tipo != 1 && tipo != 2 && tipo != 3 {
		f.Println("Insira sua categoria:")
		f.Println("Residencial;")
		f.Println("Comercial;")
		f.Println("Industrial.")
		f.Scan(&tipo)
	}
	// Escolha da categoria
	switch tipo {
	case 1:

		valor = 5 + (metros * 0.05)

	case 2:
		//Limitação do valor para cálculo excedente
		if metros <= 80 {
			valor = 500
		} else {
			valor = 500 + (metros-80)*0.25
		}
	case 3:
		if metros <= 100 {
			valor = 800
		} else {
			valor = 800 + (metros-100)*0.04
		}
	}
	f.Printf(
		"\nConta: %d\nValor: R$ %.2f", conta, valor)
}
