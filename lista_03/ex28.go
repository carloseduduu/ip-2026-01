package main

import (
	f "fmt"
	"math"
)

func main() {
	var (
		soma        float64
		denominador = 1.0
	)

	for i := 1; i <= 51; i++ {
		if i%2 != 0 {
			soma = soma + (1 / math.Pow(denominador, 3))
		} else {
			soma = soma - (1 / math.Pow(denominador, 3))
		}
		denominador += 2
	}

	valorpi := math.Cbrt(soma * 32)
	f.Printf("O valor Pi com 51 termos é: %.10f\n", valorpi)
}
