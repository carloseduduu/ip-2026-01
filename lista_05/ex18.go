package main

import f "fmt"

func main() {
	var (
		vetor [10]int
	)

	for l := 0; l < len(vetor); l++ {
		f.Scan(&vetor[l])
		lido := vetor[l]
		j := l - 1

		for j >= 0 && vetor[j] > lido {
			vetor[j+1] = vetor[j]
			j--
		}
		vetor[j+1] = lido

	}

	f.Printf("Vetor ordenado: %d", vetor)
}
