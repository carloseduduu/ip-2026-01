package main

import f "fmt"

func main() {
	var (
		n      int
		n1, n2 = 0, 1
	)

	f.Print("Insira o valor de N termos: ")
	f.Scan(&n)

	f.Printf("%d - %d - ", n1, n2)

	for i := 2; i <= n; i++ {
		atual := n1 + n2
		n1 = n2
		n2 = atual
		f.Printf("%d - ", atual)
	}
}
