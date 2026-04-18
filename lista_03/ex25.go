package main

import f "fmt"

func main() {
	var (
		divisor = 225
		razao   = 29
		soma    float64
	)

	for valor := 1.0; valor <= 16384.0; valor *= 2.0 {

		if divisor%2 == 0 {
			soma -= (valor / float64(divisor))
		} else {
			soma += (valor / float64(divisor))
		}

		divisor = divisor - razao
		razao = razao - 2
		if divisor == 0 {
			break
		}

	}
	f.Println(soma)
}
