package main

import f "fmt"

func main() {
	var n1 int
	f.Print("Insira um número inteiro de 3 dígitos: ")
	f.Scan(&n1)
	if n1 < 99 && n1 > 1000 {
		f.Print("O número digitado não possui 3 casas")
	} else {

		d := (n1 / 10) % 10
		f.Printf("O número %d, tem %d dezenas.", n1, d)
	}
}
