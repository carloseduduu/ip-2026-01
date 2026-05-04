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

	for i := 0; i < len(vetor)-1; i++ {

		if vetor[i] < vetor[i+1] {

		}

	}

	f.Printf("O menor elemento do vetor é %d, e sua posição dentro do vetor é %d", menor, posicao)

}
