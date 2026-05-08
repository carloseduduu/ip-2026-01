package main

import f "fmt"

func main() {
	var (
		cod   int
		vetor [10]float64
	)

	f.Scan(&cod)

	switch cod {
	case 0:
	case 1:
		for i, _ := range vetor {
			f.Scan(&vetor[i])
		}
		f.Printf("Vetor direto: %.2f\n", vetor)
	case 2:
		for i, _ := range vetor {
			f.Scan(&vetor[i])
		}
		f.Print("[ ")
		for i := len(vetor) - 1; i >= 0; i-- {
			f.Print(vetor[i])
			if i != 0 {
				f.Print(" ")
			}
		}
		f.Print(" ]")
	}

}
