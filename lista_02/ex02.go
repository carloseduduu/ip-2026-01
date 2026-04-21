package main

import f "fmt"

func main() {
	var valor int
	f.Scan(&valor)
	if valor < 0 {
		f.Println("Número negativo")
	} else if valor > 0 {
		f.Println("Número positivo")
	} else {
		f.Println("Número nulo")
	}
}
