package main

import f "fmt"

func main() {
	const numNotas int = 3
	var notas [numNotas]float64
	var soma float64 = 0
	for i := 0; i < numNotas; i++ {
		f.Print("Insira nota: ")
		f.Scan(&notas[i])
		soma = (soma + notas[i])
	}
	media := soma / 3.0

	f.Printf("Media = %.2f\n", media)
	for i, v := range notas {
		f.Printf("%d nota: %.2f\n", i+1, v)
	}
}
