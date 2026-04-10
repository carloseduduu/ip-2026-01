package main

import f "fmt"

func main() {
	var soma int = 0
	valor := 50
	for valor = 50; valor <= 70; valor++ {
		if valor%2 == 0 {
			soma = soma + valor
			f.Println(soma)
		}

	}
}
