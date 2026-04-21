package main

import f "fmt"

func main() {
	var (
		valores []float64
		n       float64
	)

	f.Print("Insira os valores para somar, digite -1 para parar")

	for {
		f.Scan(&n)
		if n < 0 {
			break
		}
		valores = append(valores, n)
	}

	f.Printf("A soma dos valores é: %.f", somaArray(valores))

}

func somaArray(arr []float64) float64 {
	// Caso base
	if len(arr) == 0 {
		return 0.0
	}
	// Passo recursivo: pega o 1º elemento e soma com a recursão do resto
	return arr[0] + somaArray(arr[1:])
}
