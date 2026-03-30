package main

import f "fmt"

func main() {
	var n1, i float64
	var soma float64
	f.Scan(&n1)
	for i = 1; i <= n1; i++ {
		soma += (1 / i)

	}
	f.Printf("%.6f\n", soma)
}
