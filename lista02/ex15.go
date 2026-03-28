package main

import f "fmt"

func main() {
	var dia, mes, ano int
	f.Scan(&dia, &mes, &ano)

	if mes == 1 {
		mesext := "Janeiro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 2 {
		mesext := "Fevereiro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 3 {
		mesext := "Março"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 4 {
		mesext := "Abril"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 5 {
		mesext := "Maio"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 6 {
		mesext := "Junho"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 7 {
		mesext := "Julho"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 8 {
		mesext := "Agosto"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 9 {
		mesext := "Setembro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 10 {
		mesext := "Outubro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 11 {
		mesext := "Novembro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	} else if mes == 12 {
		mesext := "Dezembro"
		f.Printf("%d de %s de %d", dia, mesext, ano)
	}
}
