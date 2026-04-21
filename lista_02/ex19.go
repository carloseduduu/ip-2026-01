package main

import (
	f "fmt"
	"math"
)

func main() {
	var tipo int
	var raio, altura float64
	var pi float64 = 3.14
	f.Println("Qual a figura desejada para cálculo do volume e àrea da superfície?")
	f.Println("1 - Cone")
	f.Println("2 - Cilindro")
	f.Println("3 - Esfera")
	f.Scan(&tipo)

	switch tipo {
	case 1:
		f.Println("Insira respectivamente o raio e a altura")
		f.Scan(&raio, &altura)
		volume := (pi * (raio * raio) * altura) / 3
		raiz := math.Sqrt(raio*raio + altura*altura)
		area := pi * raio * (raiz)
		f.Printf("O volume do cone é: %.2f\nA área do cone é %.2f.", volume, area)
	case 2:
		f.Println("Insira respectivamente o raio e a altura")
		f.Scan(&raio, &altura)
		volume := pi * raio * raio * altura
		area := 2 * pi * raio * altura
		f.Printf("O volume do cilindro é: %.2f\nA área do cilindro é %.2f.", volume, area)
	case 3:
		f.Println("Insira o raio da esfera")
		f.Scan(&raio)
		volume := (4 / 3) * pi * (raio * raio * raio)
		area := 4 * pi * raio * raio
		f.Printf("O volume da esfera é: %.2f\nA área da esfera é %.2f.", volume, area)
	}
}
