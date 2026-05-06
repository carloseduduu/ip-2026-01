package main

import f "fmt"

func main() {
	var (
		idade                    [50]int
		maiorContagem, idademoda int
	)
	busca := make(map[int]int)

	for i, _ := range idade {
		f.Scan(&idade[i])
		busca[idade[i]]++
	}

	for idade, moda := range busca {
		if moda > maiorContagem {
			maiorContagem = moda
			idademoda = idade
		}
	}

	f.Printf("Idade %d, Moda: %d\n", idademoda, maiorContagem)

}
