package main

import (
	f "fmt"
	"slices"
)

func main() {

	var (
		tamanho   int   //TAMANHO N
		num       int   //INTEIRO K
		A         []int //lista de numeros
		valor     int
		intervalo []int
		soma      int
		chave     int
	)

	f.Scan(&tamanho, &num)

	for i := 1; i <= tamanho; i++ { // lEITURA Da qntd DOS NÚMEROS
		f.Scan(&valor)
		A = append(A, valor)
	}

	for i := 1; i <= num; i++ { //Encontrar todos os termos de 1 a K
		intervalo = append(intervalo, i)
	}

	slices.Sort(A)

	for i := 0; i <= tamanho; i++ {

		chave = intervalo[i]

		for _, valor := range A {

			if chave == valor {
				continue
			} else {
				soma += chave
			}

		}
	}

	f.Println(soma)

}
