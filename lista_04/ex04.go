package main

import (
	"fmt"
	f "fmt"
)

func main() {

	decimal := 0
	f.Scan(&decimal)
	fmt.Printf("O número %d em binário é %s\n", decimal, decimalParaBinario(decimal))
}

func decimalParaBinario(num int) string {
	if num == 0 {
		return "0"
	}
	if num == 1 {
		return "1"
	}
	resto := fmt.Sprintf("%d", num%2)
	return decimalParaBinario(num/2) + resto
}
