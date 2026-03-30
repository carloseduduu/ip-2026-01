package main

import f "fmt"

func main() {
	var valor int
	f.Scan(&valor)
	if valor > 5 && valor < 2000 {
		for i := 2; i <= valor; i = i + 2 {
			quadrado := i * i
			f.Printf("%d^2 = %d\n", i, quadrado)
		}

	}

}
