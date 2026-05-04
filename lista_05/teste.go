package main

import "fmt"

func main() {
	vetorA := [10]int{5, 8, 5, 10, 5, 2, 8, 3, 3, 5}
	var processados []int

	for i := 0; i < len(vetorA); i++ {
		elemento := vetorA[i]

		// Verifica se já contamos este número antes
		jaFoiContado := false
		for _, p := range processados {
			if p == elemento {
				jaFoiContado = true
				break
			}
		}

		if !jaFoiContado {
			contador := 0
			for j := 0; j < len(vetorA); j++ {
				if vetorA[j] == elemento {
					contador++
				}
			}

			if contador > 1 {
				fmt.Printf("O valor %d se repete %d vezes\n", elemento, contador)
			}
			// Adiciona aos processados para não repetir a impressão
			processados = append(processados, elemento)
		}
	}
}
