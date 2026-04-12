package main

import (
	f "fmt"
)

func main() {
	type pessoa struct {
		idade  int
		altura float64
		peso   float64
	}

	var (
		p            pessoa
		cadastro     []pessoa
		decisao      int     = 1
		somaaltura   float64 = 0
		total                = 0.0
		media_altura float64
		peso         float64
		porcent      float64
		count50      = 0
	)

	for i := 1; decisao == 1; i++ {

		f.Print("Insira sua idade: ")
		f.Scan(&p.idade)
		f.Print("Insira sua altura em cm: ")
		f.Scan(&p.altura)
		f.Print("Insira seu peso: ")
		f.Scan(&p.peso)

		//Validação

		if p.idade <= 0 || p.altura <= 0 || p.peso <= 0 {
			f.Print("Valores inválidos\n")
			break
		}

		// Salvar cadastro na slice
		cadastro = append(cadastro, p)

		//Confimação de novo cadastro
		f.Println("Digite: \n1 - Novo usuário\n2 - Finalizar cadastro")
		f.Scan(&decisao)

		if decisao != 1 {
			f.Println("Cadastro finalizado!")
			f.Printf("Foi/foram cadastrado(s) %v usuário(s)!\n", len(cadastro))
			break
		}
	}

	// Processamento dos dados

	for _, p := range cadastro {
		// Contador idade > 50

		if p.idade > 50 {
			count50 = count50 + 1
		}

		//média das alturas 10 < idade < 20

		if p.idade >= 10 && p.idade <= 20 {
			somaaltura = somaaltura + p.altura
			total = total + 1
			media_altura = somaaltura / total
		}

		//Porcentagem de pessoas com peso inferior a 40Kg dentre todas

		if p.peso < 40 {
			peso = peso + 1
		}
		tamanho := len(cadastro)
		porcent = (peso / float64(tamanho)) * 100
	}

	f.Printf("A quantidade de pessoas com idade superior a 50 anos é %v\n", count50)

	if total > 0 {
		f.Printf("A média da altura de pessoas com idade entre 10 e 20 anos é %.2f\n", media_altura)
	} else {
		f.Printf("Nenhuma pessoa cadastrada na faixa de 10 a 20 anos\n")
	}

	f.Printf("A porcentagem de pessoas com menos de 40Kg é %.2f%%\n", porcent)
}
