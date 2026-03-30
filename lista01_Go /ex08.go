package main

import f "fmt"

func main() {
	var raio, altura float64
	const Pi float64 = 3.14159
	f.Scan(&raio)
	f.Scan(&altura)
	m := 100
	at := (2.0 * (Pi * raio * raio)) + (2.0 * Pi * raio * altura)
	custo := at * float64(m)
	f.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
