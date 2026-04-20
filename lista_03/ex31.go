package main

import f "fmt"

func main() {
	var (
		soma uint64 = 0
		Qntd uint64 = 1
	)

	for i := 1; i <= 64; i++ {
		f.Printf("Quadro: %d, Qntd: %d, Soma:%d\n", i, Qntd, soma)
		soma = soma + Qntd
		Qntd = Qntd * 2

	}
}
