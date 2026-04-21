package main

import f "fmt"

func main() {

	var teste, pop, ger, arq, cad, ingresso float64

	const (
		valor_popular      = 1
		valor_ger          = 5
		valor_arquibancada = 10
		valor_cadeiras     = 20
	)
	var renda_lista []float64

	f.Println("Quantos testes?")
	f.Scan(&teste)

	for i := 0; i < int(teste); i++ {

		f.Scan(&ingresso, &pop, &ger, &arq, &cad)
		renda := ((ingresso * (pop / 100.0)) * valor_popular) + ((ingresso * (ger / 100.0)) * valor_ger) + ((ingresso * (arq / 100.0)) * valor_arquibancada) + ((ingresso * (cad / 100.0)) * valor_cadeiras)
		renda_lista = append(renda_lista, renda)
		f.Print("\n")
	}

	for i, v := range renda_lista {
		f.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, v)
	}

}
