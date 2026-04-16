package main

import (
	f "fmt"
)

func main() {
	var (
		base, exp int
	)
	f.Print("Insira base e potência")
	f.Scan(&base, &exp)

}

func potência(base, expoente int) int {
	for i := 2; i < expoente; i++ {
		resultado := base * resultado
		base = resultado
	}
}
