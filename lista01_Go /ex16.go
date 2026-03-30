package main

import f "fmt"

func main() {
	var salario float64
	f.Scan(&salario)
	switch {
	case salario <= 300:
		salario += (salario * 50.0) / 100.0
		f.Printf("SALARIO COM REAJUSTE = %.2f\n", salario)
	case salario > 300:
		salario += (salario * 30.0 / 100.0)
		f.Printf("SALARIO COM REAJUSTE = %.2f\n", salario)
	}
}
