package main

import f "fmt"

func main() {
	var (
		vetor [10]int
	)

	for i := 0; i < len(vetor); i++ {
		f.Scan(&vetor[i])
	}

	for i := 0; i < len(vetor); i++ {
		ePrimo := true
		for j := 2; j <= vetor[i]/2; j++ {

			if vetor[i]%j == 0 {
				ePrimo = false
				break
			}

		}
		
		if ePrimo {
			f.Printf("Valor: %d, Posição: %d\n", vetor[i], i)
		}
	}
}
