package main

import f "fmt"

func main() {
	var (
		vetor              [10]int
		somapar, somaimpar int
		pares, impar       []int
	)

	for i := 0; i < len(vetor); i++ {
		f.Scan(&vetor[i])
	}

	for i := 0; i < len(vetor); i++ {
		if vetor[i]%2 == 0 {
			pares = append(pares, vetor[i])
			somapar = somapar + vetor[i]
		} else {
			impar = append(impar, vetor[i])
			somaimpar = somaimpar + vetor[i]
		}
	}

	f.Printf("Números pares %d, Soma pares %d \nNúmeros Ímpares %d, Soma Impar %d\n", pares, somapar, impar, somaimpar)
}
