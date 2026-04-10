package main

import f "fmt"

func main() {

	var n1, n2, n3 int
	f.Scan(&n1, &n2, &n3)
	if n1 > 9 || n2 > 9 || n3 > 9 {
		f.Println("DIGITO INVALIDO")
	} else {
		valor := (n1 * 100) + (n2 * 10) + n3
		quadrado := valor * valor
		f.Printf("%d,%d", valor, quadrado)
	}
}
