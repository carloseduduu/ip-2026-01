package main

import f "fmt"

func main() {

	var n1, n2, n3 float64
	f.Println("Insira suas 3 notas")
	f.Scan(&n1, &n2, &n3)
	media := (n1 + n2 + n3) / 3.0
	switch {

	case media >= 6:
		f.Printf("MEDIA = %.2f\n", media)
		f.Print("APROVADO\n")
	case media < 6:
		f.Printf("MEDIA = %.2f\n", media)
		f.Println("REPROVADO\n")
	}
}
