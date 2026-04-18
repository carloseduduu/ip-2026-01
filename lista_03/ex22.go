package main

import f "fmt"

func main() {

	var divisor1, divisor2, soma = 37.0, 38.0, 0.0

	for i := 1; i <= 37; i++ {

		soma += ((divisor1 * divisor2) / float64(i))

		if divisor1 != 1 {
			divisor1--
		}

		if divisor2 != 2 {
			divisor2--
		}
	}

	f.Printf("Soma: %.2f\n", soma)
}
