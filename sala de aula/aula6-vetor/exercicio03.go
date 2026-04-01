package main

import f "fmt"

func main() {
	const limite int = 10
	var numeros [limite]float64
	for i := 0; i < limite; i++ {
		f.Scan(&numeros[i])
	}
	for i := len(numeros); i > 0; i-- {
		f.Printf("Número na posição %d = %.2f\n", i, numeros[i-1])
	}
}
