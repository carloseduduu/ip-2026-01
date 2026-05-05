package main

import (
	f "fmt"
	"math"
	"math/rand"
)

func main() {
	var (
		valor [100]int
		s     float64
	)
	for i := 0; i < len(valor); i++ {
		valor[i] = rand.Intn(15)
	}

	f.Println(valor)
	for i, f := 0, 99; i < f; i, f = i+1, f-1 { // i início, f fim
		s += math.Pow(float64(valor[i]-valor[f]), 3)
	}

	f.Println(s)
}
