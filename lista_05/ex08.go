package main

import (
	f "fmt"
	"math"
)

func main() {
	var (
		raiz  [15]float64
		valor int
	)

	for i, _ := range raiz {
		f.Scan(&valor)
		if valor < 0 {
			raiz[i] = -1
		} else {
			raiz[i] = (math.Sqrt(float64(valor)))
		}
	}

	f.Println(raiz)
}
