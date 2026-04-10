package main

import f "fmt"

func main() {
	var farenheit, polegadas float64
	f.Scan(&farenheit, &polegadas)
	celsius := (5.0*farenheit - 160.0) / 9.0
	mm := polegadas * 25.4
	f.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f\n", celsius, mm)
}
