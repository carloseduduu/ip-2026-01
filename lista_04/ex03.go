package main

import (
	f "fmt"
)

func main() {
	var (
		valores []int
		n       int
	)

	f.Print("Insira os valores para somar, digite -1 para parar\n")

	for {
		f.Scan(&n)
		if n < 0 {
			break
		}
		valores = append(valores, n)
	}

	inverteArray(valores, len(valores))
	f.Printf("Os valores inverso é: %v\n", valores)
}

func inverteArray(arr []int, n int) {
	if n <= 1 {
		return
	}
	arr[0], arr[n-1] = arr[n-1], arr[0]
	inverteArray(arr[1:n-1], n-2)
}
