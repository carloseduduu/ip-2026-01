package main

import f "fmt"

func main() {
	var (
		valor float64
	)

	for i := 1; i <= 99; i++ {
		valor = float64(i/i) + float64(i+2/i)
	}
	f.Print(valor)
}
