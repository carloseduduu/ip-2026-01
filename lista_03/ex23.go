package main

import f "fmt"

func main() {
	var (
		ntermos, divisor, resultado = 0.0, 1000.0, 0.0
	)

	f.Scan(&ntermos)

	for i := 1; float64(i) <= ntermos; i++ {

		if i%2 != 0 {
			resultado = resultado + (divisor / float64(i))
		} else {
			resultado = resultado - (divisor / float64(i))
		}

		divisor -= 3
	}

	f.Printf("O resultado é: %.2f\n", resultado)
}
