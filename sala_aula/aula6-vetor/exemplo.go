package main

import f "fmt"

func main() {
	const numNotas int = 3
	var (
		nota [numNotas]float64
		soma float64 = 0
	)
	for i := 0; i < numNotas; i++ {
		f.Printf("Informe a nota %d: ", i)
		f.Scan(&nota[i])
		soma += nota[i]
	}
	for i, v := range nota {
		f.Print("Nota %d = %f\n", i, v)
	}
	f.Printf("Soma das notas %f: ", soma)
}
