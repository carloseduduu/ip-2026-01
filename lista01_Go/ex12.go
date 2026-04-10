package main

import f "fmt"

func main() {
	var horas int
	f.Scan(&horas)
	resto := (horas % 3)
	valor := 10.0*(horas/3.0) + 5.0*(resto)
	f.Printf("O VALOR A PAGAR E = %d", valor)
}
