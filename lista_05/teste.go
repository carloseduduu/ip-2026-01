package main

import f "fmt"

func main() {
	var vetor [5]int

	for i := 0; i < cap(vetor); i++ {
		f.Scan(&vetor[i])
	}
	menorValor := vetor[0]
	posicao := 0

	for i, v := range vetor {
		if v < menorValor {
			menorValor = v
			posicao = i
		}
	}

	f.Printf("Valor posição 0: %d\nValor posição 4: %d\n", vetor[0], vetor[4])
	f.Printf("Menor valor: %d\n", menorValor)
	f.Printf("Posição menor valor: %d", posicao)
}
