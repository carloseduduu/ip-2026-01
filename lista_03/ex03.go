package main

import f "fmt"

func main() {
	var carlos, joao, i float64
	f.Print("Insira o salário de Carlos: ")
	f.Scan(&carlos)
	joao = carlos / 3
	f.Printf("O salário de João é R$ %v\n", joao)

	for i = 1; joao <= carlos; i++ {
		carlos += (carlos * 2 / 100)
		joao += (joao * 5 / 100)
		f.Printf("Carlos mês %v = R$ %.2f\nJoao mês %v = R$ %.2f\n\n", i, carlos, i, joao)
	}

	f.Printf("A quantidade  de meses necessários para que o valor pertencente a João iguale ou ultrapasse o valor pertencente a Carlos foi de %.f meses", i-1)
}
