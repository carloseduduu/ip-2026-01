package main

import f "fmt"

func main() {
	var (
	//linha, coluna int
	)

	for i := 0.0; i <= 6.3; i += 0.1 {

		f.Printf("(Sen%.1f=%.1f)", i, sen(i))
		f.Print("\n")
	}

}

func sen(ang float64) float64 {

	seno := ang - ((potencia(ang, 3)) / 6) + (potencia(ang, 5))/120 - (potencia(ang, 7))/5040
	return seno
}

func potencia(base, expoente float64) float64 {
	var resultado float64 = 1
	for i := 0.0; i < expoente; i++ {

		resultado = resultado * base

	}
	return resultado
}
