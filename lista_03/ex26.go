package main

import f "fmt"

func main() {
	var (
		soma  float64
		termo = 0.0
	)
	for i := 100; i >= 80; i-- {

		soma += float64(i) / fat26(termo)
		f.Printf("Termo: %v Soma: %.2f Índice: %v\n", termo, soma, i)
		termo++
	}
}

func fat26(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * fat26(n-1)
}
