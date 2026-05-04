package main

import f "fmt"

func main() {
	var (
		vetor          [10]int
		menor, posicao int
	)

	for i := 0; i < len(vetor); i++ {
		f.Scan(&vetor[i])
	}
	menor = vetor[0]
	posicao = 0

	for i, v := range vetor {
		if v < menor {
			menor = v
			posicao = i
		}
	}

	f.Printf("O menor elemento do vetor é %d, e sua posição dentro do vetor é %d", menor, posicao)

}
