package main

import (
	f "fmt"
)

func main() {
	var (
		n1, n2  int
		números []int
	)

	f.Print("Determine o início e o fim do intervalor (N1 e N2)")
	f.Scan(&n1, &n2)

	for i := n1; i <= n2; i++ {
		if i < 2 {
			continue
		}

		primo := true

		for divisor := 2; divisor < i; divisor++ {
			if i%divisor == 0 {
				primo = false
				break
			}
		}

		if primo == true {
			números = append(números, i)
		}
	}

	f.Println("\nNúmeros primos encontrados:", números)
}
