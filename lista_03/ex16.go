package main

import f "fmt"

func main() {
	var (
		n1, n2, n int
	)
	f.Print("Insira n1, n2 e N termos: ")
	f.Scan(&n1, &n2, &n)

	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			valor := n1 + n2
			n1 = n2
			n2 = valor
			f.Printf("O %vº termo é: %v\n", i, valor)
		} else {
			valor := n1 - n2
			n1 = n2
			n2 = valor
			f.Printf("O %vº termo é: %v\n", i, valor)
		}

	}
}
