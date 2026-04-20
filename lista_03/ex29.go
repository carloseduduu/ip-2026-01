package main

import f "fmt"

func main() {
	var (
		soma, n int
	)

	f.Scan(&n)
	for i := 1; i <= n; i++ {
		soma = soma + i
	}
	f.Println(soma)
}
