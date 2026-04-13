package main

import f "fmt"

func main() {
	var (
		n, ant, termo int
	)

	f.Print("Insira o valor de N termos: ")
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		ant = termo
		termo = ant + i
		f.Printf("%d-", termo)
	}
}
