package main

import (
	f "fmt"
	"math"
)

func main() {
	var (
		valor     string
		lista     []int
		resultado float64
	)

	f.Scan(&valor)
	tamanho := len(valor) - 1            // Se são três valores, tamanho será 3-1 = 2 ou seja posições 0, 1, 2.
	for i := 0; i <= len(valor)-1; i++ { // Contador i é o expoente do 8 para cada posição percorrida.

		num := int(valor[tamanho] - '0') // Conversão de string, para int

		resultado += (float64(num) * math.Pow(8, float64(i))) // Transforma uma unidade de base 8 em unidade de base 10

		tamanho-- // Percorre a lista de trás pra frente
	}

	lista = append(lista, int(resultado))
	for i := len(lista) - 1; i >= 0; i-- {
		f.Printf("%d", lista[i])
	}

}
