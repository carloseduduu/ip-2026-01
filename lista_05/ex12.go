package main

import f "fmt"

func main() {
	var (
		notas     = [15]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		contador  int
		jaContado []float64
	)

	for i, _ := range notas {
		f.Scan(&notas[i])
	}
	f.Print("|NOTA\t|ABSOLUTA|RELATIVA|\n")
	for i := 0; i < len(notas); i++ {
		contado := false
		contador = 0
		valor := notas[i]

		for _, v := range jaContado {
			if valor == v {
				contado = true
				break
			}
		}

		if !contado {

			for i := 0; i < len(notas); i++ {

				if valor == notas[i] {

					contador++

				}

			}
			jaContado = append(jaContado, valor)
			relativa := float64(contador) / float64(len(notas))
			f.Printf("|%.f\t|%d\t |%.2f\t|\n", valor, contador, relativa)

		}

	}

}
