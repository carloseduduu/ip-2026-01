package main

import (
	"fmt"
)

func main() {
	var (
		x    float64
		soma float64
	)

	fmt.Print("Insira um número real (X): ")
	fmt.Scan(&x)

	for i := 2; i < 20; i++ {
		termo := x / fatorial(i)
		soma = x - termo
		fmt.Printf("Iteração %d: S vale %f\n", i, soma)
	}

	fmt.Printf("O resultado de S é: %f\n", soma)
}

func fatorial(n int) float64 {
	resultado := 1.0
	for i := 1; i <= n; i++ {
		resultado = resultado * float64(i)
	}
	return resultado
}
