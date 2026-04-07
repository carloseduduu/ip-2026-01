package main

import (
	f "fmt"
)

func main() {
	type pessoa struct {
		nome      string
		altura    float64
		pesoideal float64
	}
	var dados pessoa
	var limite int

	f.Print("Quantas pessoas quer adicionar?")
	f.Scan(&limite)

	pessoas := make([]pessoa, limite)

	for i := 0; i < limite; i++ {
		if dados.nome != "FIM" {
			f.Print("Insira seu nome: ")
			f.Scan(&dados.nome)
			f.Print("Insira sua altura: ")
			f.Scan(&dados.altura)
			dados.pesoideal = 72.7*dados.altura - 58.0
			pessoas[i] = dados
		}
	}

	for i := 0; i < limite; i++ {
		f.Print(pessoas[i])
	}

}
