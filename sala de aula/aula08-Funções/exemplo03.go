package main

import f "fmt"

func media(nota1, nota2, nota3 int) float64 {
	m := float64(nota1+nota2+nota3) / 3
	return m
}

func main() {
	var n1, n2, n3 int
	f.Scan(&n1, &n2, &n3)
	f.Printf("A media é: %.2f\n", media(n1, n2, n3))
}
