package main

import f "fmt"

func main() {
	var (
		tamanho, valor int
		contador       int
		lista          []int
	)

	f.Scan(&tamanho)

	for leitura := 1; leitura <= tamanho; leitura++ {
		f.Scan(&valor)

		lista = append(lista, valor)
	}

	for i := 0; i < len(lista)-2; i++ {

		if lista[i] == 1 {

			for i = i + 1; i < len(lista); i++ {
				if lista[i] == 0 {

					for i = i + 2; i < len(lista); i++ {
						if lista[i] == 0 {
							//posicao3 = lista[i]
							contador += 1
						} else {
							break
						}
					}
				} else {
					break
				}
			}

		}
	}
	f.Println(contador)
}
