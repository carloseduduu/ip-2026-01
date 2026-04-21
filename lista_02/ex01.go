package main

import f "fmt"

func main() {
	var valor int
	f.Scan(&valor)
	if valor == 0 {
		f.Print("VALOR INVÁLIDO")
	} else if valor%2 == 0 {
		f.Print("O NÚMERO É PAR")
	} else {
		f.Print("O NÚMERO É ÍMPAR")
	}
}
