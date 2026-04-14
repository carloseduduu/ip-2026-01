package main

import f "fmt"

func main() {
	var numero int
	f.Scan(&numero)
	f.Printf("O numero em fatorial é: %v\n", fatorial(numero))
}

func fatorial(n1 int) int {
	resultado := 1
	for i := 1; i <= n1; i++ {
		resultado = resultado * i
	}
	return resultado
}
