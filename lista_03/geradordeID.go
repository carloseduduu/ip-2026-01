package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed para garantir que os pesos sejam diferentes a cada execução
	rand.Seed(time.Now().UnixNano())

	fmt.Print("{")

	for i := 1; i <= 90; i++ {
		// 1. Gerar o ID: i formato com 2 dígitos (01, 02...)
		// 2. Gerar o Peso: entre 400.0 e 900.0 (exemplo de peso de abate)
		pesoInteiro := rand.Intn(500) + 400 // Gera de 400 a 899
		pesoDecimal := rand.Intn(10)        // Gera de 0 a 9

		// Formatação: %02d (2 dígitos com zero à esquerda)
		// %03d (3 dígitos inteiros para o peso)
		// %d (1 dígito decimal)
		if i == 90 {
			fmt.Printf("\"%02d%03d%d\"}", i, pesoInteiro, pesoDecimal)
			break
		}
		fmt.Printf("\"%02d%03d%d\", ", i, pesoInteiro, pesoDecimal)
	}

}
