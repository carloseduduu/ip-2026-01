package main

import f "fmt"

func main() {
	var n1 float64
	f.Print("Insira o valor de x para obter o valor de f(x), dado que:f(x) = 8 / (2-x): ")
	f.Scan(&n1)
	y := 8 / (2 - (n1))
	f.Printf("O Valor de f(%.1f) é %.2f", n1, y)
}
