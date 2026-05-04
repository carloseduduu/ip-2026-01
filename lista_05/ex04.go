package main

import f "fmt"

func main() {
	var (
		vetorA   [10]int
		contador int
	)

	type lista struct {
		valor, qnt []int
	}

	var l lista

	for i := 0; i < len(vetorA); i++ { //Leitura dos 10 valores
		f.Scan(&vetorA[i])
	}

	for i := 0; i < len(vetorA); i++ { // Seleciona cada índice
		contador = 0
		elemento := vetorA[i]

		for i := 0; i < len(vetorA); i++ { // Compara o índice com todos os elementos

			if elemento == vetorA[i] { // se o elemento do índice for igual ao item dentro do vetor soma 1 ao contador
				contador += 1
			}
		}

		if contador != 0 {
			l.qnt = append(l.qnt, contador)
			l.valor = append(l.valor, vetorA[i])
		}
	}

	for i := 0; i < len(l.valor); i++ {
		f.Printf("Valor %d repete %d vezes \n", l.valor[i], l.qnt[i])
	}
}
