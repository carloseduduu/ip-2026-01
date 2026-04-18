package main

import (
	f "fmt"
	"math"
)

func main() {

	var (
		x, soma, cosseno float64
		divisor          float64 = 2
	)

	f.Scan(&x)

	for i := 1; i < 20; i++ { //RODA 20X COM I SUBINDO DE 2 EM 2

		if i%2 == 0 {
			soma = soma - (math.Pow(x, divisor))/fat27(divisor)
		} else {
			soma = soma + (math.Pow(x, divisor))/fat27(divisor)
		}
		divisor += 2

	}

	cosseno = 1 - soma
	diferença := math.Abs(cosseno - math.Cos(x))
	f.Printf("Item A: %f, Item B: %f, Cos(x) = %f\n", cosseno, diferença, math.Cos(x))
}

func fat27(n float64) float64 { //Retorna o fatorial de um número
	if n == 0 {
		return 1
	}
	return n * fat27(n-1)
}
