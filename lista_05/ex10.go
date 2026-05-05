package main

import f "fmt"

func main() {
	//Fibonacci
	var (
		vetor = [50]int{1, 1}
	)
	f.Println(vetor[0])

	for i := 2; i < 50; i++ {
		vetor[i] = vetor[i-1] + vetor[i-2]
	}

	f.Println(vetor)
}
