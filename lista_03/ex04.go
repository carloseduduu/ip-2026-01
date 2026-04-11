package main

import (
	f "fmt"
)

func main() {
	var valor int

	for {
		f.Print("Insira um valor inteiro positivo (Digite 0 para finalizar): ")
		f.Scan(&valor)

		if valor <= 0 {
			break
		}
		i := 1
		for i <= valor {
			if i*i == valor {
				f.Printf("%d é um quadrado perfeito\n", valor)
				break
			} else if i == valor {
				f.Printf("%d não é um quadrado perfeito\n", valor)
			}
			i++
		}
	}
}
