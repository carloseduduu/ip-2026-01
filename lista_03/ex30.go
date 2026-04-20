package main

import (
	f "fmt"
	"math"
)

func main() {

	for i := 0.0; i <= 20; i += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * (math.Pow(i, 3))
		f.Printf("Raio: %.1f, Volume: %.2f\n", i, volume)
	}
}
