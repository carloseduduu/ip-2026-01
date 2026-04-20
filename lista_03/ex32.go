package main

import f "fmt"

func main() {
	var (
		n1, n2, soma int
	)

	f.Scan(&n1, &n2)
	soma = n1
	for i := 1; i < n2; i++ {
		soma = soma + n1
	}
	f.Printf("Resultado %dx%d = %d\n", n1, n2, soma)
}
