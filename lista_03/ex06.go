package main

import f "fmt"

func main() {
	var (
		valor     int
		resultado int
	)

	f.Print("Insira um valor inteiro positivo: ")
	f.Scan(&valor)

	for n := 1; n*(n+1)*(n+2) <= valor; n++ {
		resultado = n * (n + 1) * (n + 2)

	}

	if resultado == valor {
		f.Printf("O número %d é triangular!\n", valor)
	} else {
		f.Printf("O número %d NÃO é triangular.\n", valor)
	}
}
