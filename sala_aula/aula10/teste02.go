package main

import f "fmt"

func main() {
	var p1, p2, media float64

	f.Print("Insira sua pontuação 1: ")
	f.Scan(&p1)
	f.Print("Insira sua pontuação 2: ")
	f.Scan(&p2)
	media = ((p1 * 2.0) + (p2 * 3.0)) / 5.0

	if p1 >= 0.0 && p1 <= 10.0 && p2 >= 0.0 && p2 <= 10.0 {

		switch {
		case media >= 7.0:
			f.Println("Vitorioso")

		case media < 3.0:
			f.Println("Derrotado")

		case media >= 3.0 && media < 7.0:
			f.Println("Julgamento Final")
		}
	} else {
		f.Println("Valores inválidos")
	}

}
