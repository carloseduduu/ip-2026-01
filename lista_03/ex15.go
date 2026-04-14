package main

import f "fmt"

func main() {
	var (
		n            int
		razao, valor = 3, 1
	)

	f.Print("Insira o valor N termo: ")
	f.Scan(&n)

	for i := 1; i <= n; i++ {

		f.Printf("Valor: %d\n", valor) // Exibe o valor antes de calcular

		valor2 := valor + razao
		razao = razao + 2
		valor = valor2

	}
}
