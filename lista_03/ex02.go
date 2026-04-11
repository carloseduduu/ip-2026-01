package main

import f "fmt"

func main() {
	var soma int = 0
	var i int = 0
	valor := 50
	for valor = 50; valor <= 70; valor++ {
		if valor%2 == 0 {
			soma = soma + valor
			i++
			f.Printf("A soma na posição %v é %v\n", i, soma)
		}

	}
	media := soma / i
	f.Printf("A média dos valores pares entre 50 e 70 é %v\n", media)
}
