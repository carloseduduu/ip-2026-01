package main

import f "fmt"

func main() {
	var (
		valor, ingresso, maiorlucro, ingressomaior float64 = 0, 130, 0, 0
	)

	for valor = 6.0; valor >= 1.0; valor -= 0.6 {

		lucro := (ingresso * valor) - 300
		if lucro > maiorlucro {
			maiorlucro = lucro
			ingressomaior = ingresso
		}
		f.Printf("Ingressos: %.f Valor: %.2f Lucro: %.f\n", ingresso, valor, lucro)
		ingresso += 30.0
	}
	f.Printf("O maior lucro foi de R$ %.2f. Com %.f ingressos vendidos.", maiorlucro, ingressomaior)
}
