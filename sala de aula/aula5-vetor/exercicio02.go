package main

import f "fmt"

func main() {
	const tamanho int = 5
	var numeros [tamanho]float64
	var soma float64 = 0
	for i := 0; i < tamanho; i++ {
		f.Scan(&numeros[i])
		soma += numeros[i]
	}
	f.Printf("A soma é = %.2f", soma)
}
