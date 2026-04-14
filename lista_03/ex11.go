package main

import f "fmt"

func main() {
	var numero int
	f.Print("Insira um valor para calcular fatorial: ")
	f.Scan(&numero)
	f.Printf("O numero em fatorial é: %v\n", fatorial2(numero))
}

func fatorial2(n1 int) int {
	resultado := 1
	for i := 1; i <= n1; i++ {
		resultado = resultado * i
	}
	return resultado
}
