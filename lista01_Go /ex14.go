package main

import (
	f "fmt"
	"math"
)

func main() {
	var altura, aresta float64
	f.Scan(&altura, &aresta)
	raiz := math.Sqrt(3)
	volume := (1.0 / 3.0) * ((3 * aresta * aresta * raiz) / 2) * altura
	f.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", volume)
}
