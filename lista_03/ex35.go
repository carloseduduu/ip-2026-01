package main

import f "fmt"

func main() {
	var (
		valor int
		lista []int
	)

	f.Scan(&valor)

	for {
		resto := valor % 2
		lista = append(lista, resto)
		valor = valor / 2

		if valor == 0 {
			break
		}
	}

	for i := len(lista) - 1; i >= 0; i-- {
		f.Print(lista[i])
	}

	f.Print("\n")
}
