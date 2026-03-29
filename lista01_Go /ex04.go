package main

import f "fmt"

func main() {
	var salario, kw float64

	f.Scan(&salario, &kw)
	valorkw := (salario * 0.7) / 100.0
	gasto := kw * valorkw

	f.Printf("Custo por kW: R$ %.2f\n", valorkw)
	f.Printf("Custo do consumo: R$ %.2f\n", gasto)
	gasto += -gasto * 0.1
	f.Printf("Custo com desconto: R$ %.2f\n", gasto)
}
