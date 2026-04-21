package main

import f "fmt"

func main() {
	var conta, consumo float64
	var tipo string
	f.Scanf("%f%f%s", &conta, &consumo, &tipo)

	switch tipo {
	case "R":
		valor := 5 + (consumo * 0.05)
		f.Printf("CONTA = %.f\nVALOR DA CONTA = %.2f\n", conta, valor)
	case "C":
		valor := 500 + (consumo-80)*0.25
		f.Printf("CONTA = %.f\nVALOR DA CONTA = %.2f\n", conta, valor)
	case "I":
		valor := 800 + (consumo-100)*0.04
		f.Printf("CONTA = %.f\nVALOR DA CONTA = %.2f\n", conta, valor)
	}
}
