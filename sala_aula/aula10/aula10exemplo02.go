package main

import (
	f "fmt"
)

type Pessoa struct {
	nome      string
	idade     int
	profissao string
	salario   int
}

func main() {
	var pes1 Pessoa
	// Especificacao da pes1
	/*pes1.nome = "Monica"
	pes1.idade = 35
	pes1.profissao = "Engenheiro de Software"
	pes1.salario = 7000*/
	f.Scan(&pes1.nome)
	f.Scan(&pes1.idade)
	f.Scan(&pes1.profissao)
	f.Scan(&pes1.salario)
	// Imprimir as informacoes da pessoa1
	f.Println("nome: ", pes1.nome)
	f.Println("idade: ", pes1.idade)
	f.Println("profissao: ", pes1.profissao)
	f.Println("salario: ", pes1.salario)
}
