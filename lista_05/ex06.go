package main

import f "fmt"

func main() {
	var (
		vetor [100]int
	)

	for j, i := 0, 100; i >= 1; j, i = j+1, i-1 { //Duas variáveis dentro de um mesmo for
		vetor[j] = i
	}

	for j := 0; j < cap(vetor); j++ {
		f.Printf("Elemento: %d\n", vetor[j])
	}
}
