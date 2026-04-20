package main

import f "fmt"

func main() {
	var (
		valor int
		lista []int
	)

	f.Scan(&valor)

	for {
		resto := valor % 16
		lista = append(lista, resto)
		valor = valor / 2

		if valor == 0 {
			break
		}
	}

	f.Print("\n")
}
