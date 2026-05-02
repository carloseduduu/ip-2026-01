package main

import f "fmt"

func main() {
	var (
		vetor [10]int
		valor int
	)
	type maior50 struct {
		posicao int
		valor   int
	}

	for i := 0; i < len(vetor); i++ {
		f.Scan(&valor)
		vetor[i] = valor
	}

	var (
		m     maior50
		lista []maior50
	)

	for m.posicao, m.valor = range vetor {

		if m.valor > 50 {
			lista = append(lista, m)
		}
	}

	for i, _ := range lista {
		f.Printf("Valor: %d, Posição: %d\n", lista[i].valor, lista[i].posicao)
	}
}
