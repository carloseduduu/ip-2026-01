package main

import f "fmt"

func main() {
	var (
		lista = []int{1, 5, 15, 20, 24, 45, 67, 76, 78, 100}
		valor int
	)

	f.Printf("Insira um valor para busca: ")
	f.Scan(&valor)

	resultado := buscabinaria(lista, valor)

	if resultado != -1 {
		f.Printf("O valor está na posição %d\n", resultado)
	} else {
		f.Print("O valor não está na lista\n")
	}
}

func buscabinaria(lista []int, item int) int {
	n := len(lista)
	e := 0
	d := n - 1

	for e <= d {
		m := (e + d) / 2
		if lista[m] == item {
			return m
		} else if lista[m] < item {
			e = m + 1
		} else if lista[m] > item {
			d = m - 1
		}
	}

	return -1
}
