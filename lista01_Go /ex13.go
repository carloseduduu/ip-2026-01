package main

import f "fmt"

func main() {
	var nota float64
	f.Scan(&nota)
	switch {
	case nota >= 9 && nota <= 10:
		conceito := "A"
		f.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
	case nota >= 7.5 && nota < 9:
		conceito := "B"
		f.Printf("NOTA = %.1f CONCEITO  = %s\n", nota, conceito)
	case nota >= 6 && nota < 7.5:
		conceito := "C"
		f.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
	case nota >= 0 && nota < 6:
		conceito := "D"
		f.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
	default:
		f.Print("NOTA INVALIDA\n")
	}

}
