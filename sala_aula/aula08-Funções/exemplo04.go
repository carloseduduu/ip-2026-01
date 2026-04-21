package main

import f "fmt"

func main() {
	var numero int
	f.Scan(&numero)
	f.Printf("O numero em fatorial é: %v\n", fatorial(numero))
}

func fatorial(n1 int) int {
	resultado := 1
	for i := n1; i > 0; i-- {
		resultado = resultado * i
	}
	return resultado
}
