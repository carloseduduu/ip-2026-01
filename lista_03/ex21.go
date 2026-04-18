package main

import (
	f "fmt"
)

func main() {
	var (
		base, exp float64
	)
	f.Print("Insira base e potência: ")
	f.Scan(&base, &exp)
	f.Printf("O valor %d elevado a %dª potência é igual a: %v\n", base, exp, potência(base, exp))

}

func potência(base, expoente float64) float64 {
	var resultado float64 = 1
	for i := 0.0; i < expoente; i++ {

		resultado = resultado * base

	}
	return resultado
}
