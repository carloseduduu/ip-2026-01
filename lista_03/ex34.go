package main

import (
	f "fmt"
)

func main() {
	var (
		n1, n2, resto, MDC int
	)
	f.Print("Insira os valores de N1 e N2: ")
	f.Scan(&n1, &n2)

	for { //Calcula o MDC com o algoritmo de Euclides
		resto = n1 % n2
		n1 = n2
		n2 = resto

		if resto == 0 {
			MDC = n1
			break
		}
	}
	f.Printf("MDC: %d", MDC)

	MMC := (n1 * n2) / MDC
	f.Printf("O MMC entre %d e %d é igual a: %d", n1, n2, MMC)
}
