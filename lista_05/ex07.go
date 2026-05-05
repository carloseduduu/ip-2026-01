package main

import f "fmt"

func main() {

	var (
		i, j int = 1, 0
	)
	vetor := make([]int, 0, 100)
	f.Println(cap(vetor), len(vetor))
	for {

		if i%2 != 0 {
			vetor = append(vetor, i)
			j++
		}

		if len(vetor) == cap(vetor) { // len: Retorna quantidade de valore iniciados, cap: retorna o total da capacidade
			break
		}
		i++
	}
	f.Println(cap(vetor), len(vetor), vetor)

}
