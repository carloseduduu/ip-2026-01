package main

import f "fmt"

func main() {
	var A, B, C float64
	f.Scan(&A)
	f.Scan(&B)
	f.Scan(&C)
	delta := (B * B) - 4.0*A*C
	f.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
