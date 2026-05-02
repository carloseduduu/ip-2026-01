package main

import f "fmt"

func main() {
	var (
		v     [10]int
		v2    [5]int
		vr    []int
		vr2   []int
		soma  int
		valor int
	)

	for i := 0; i < len(v); i++ {
		f.Scan(&valor)
		v[i] = valor

	}

	for i := 0; i < len(v2); i++ {
		f.Scan(&valor)
		v2[i] = valor
	}

	for _, v := range v {

		if v%2 == 0 {

			soma = v

			for i := 0; i < len(v2); i++ {
				soma = soma + v2[i]
			}
			vr = append(vr, soma)
		} else {

			soma = v

			for i := 0; i < len(v2); i++ {
				soma = soma + v2[i]
			}

			vr2 = append(vr2, soma)
		}
	}

	f.Printf("Primeiro vetor resultante: %d\n", vr)
	f.Printf("Segundo vetor resultante: %d\n", vr2)

}
