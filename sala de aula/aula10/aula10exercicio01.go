package main

import f "fmt"

func main() {

	type pessoa struct { // STRUCT -> type NOME struct {}
		nome      string
		altura    float64
		pesoideal float64
	}

	var (
		p     pessoa
		lista []pessoa
	) //SLICE -> NOME []TIPO

	for i := 0; i >= 0; i++ {

		f.Scan(&p.nome)

		if p.nome == "FIM" {
			break
		}

		f.Scan(&p.altura)
		p.pesoideal = p.altura / 2.57
		lista = append(lista, p)

	}

	for i := 0; i < len(lista); i++ {

		f.Printf("Nome: %v\nAltura: %v\nPeso ideal: %v\n", lista[i].nome, lista[i].altura, lista[i].pesoideal)
	}

	f.Println("FIM")
}
