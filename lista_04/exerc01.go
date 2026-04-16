package main

import f "fmt"

func main() {
	var (
		base, expoente int
	)
	f.Print("Insira valor base e expoente: ")
	f.Scan(&base, &expoente)

	f.Print(pot(base, expoente))
}

func pot(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * pot(base, exp-1)
}
