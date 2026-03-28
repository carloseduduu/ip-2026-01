package main

import (
	f "fmt"
	"strings"
)

func main() {
	var valor float64
	var decisao string
	f.Println("Qual o valor inicial do carro?")
	f.Scan(&valor)

	f.Println("Você deseja adicionar: Ar condicionado (R$ 1750,00)? [Sim][Não]")
	f.Scan(&decisao)
	decisao = strings.ToLower(decisao)

	if decisao == "sim" || decisao == "s" {
		valor += 1750
	}

	f.Println("Você deseja adicionar: Pintura metálica (R$ 800,00) [Sim][Não]")
	f.Scan(&decisao)
	decisao = strings.ToLower(decisao)

	if decisao == "sim" || decisao == "s" {
		valor += 800
	}

	f.Println("Você deseja adicionar: Vidro elétrico (R$ 1200,00) [Sim][Não]")
	f.Scan(&decisao)
	decisao = strings.ToLower(decisao)

	if decisao == "sim" || decisao == "s" {
		valor += 1200
	}

	f.Println("Você deseja adicionar: Direção hidráulica (R$ 2000,00) [Sim][Não]")
	f.Scan(&decisao)
	decisao = strings.ToLower(decisao)

	if decisao == "sim" || decisao == "s" {
		valor += 2000
	}
	f.Printf("O preço final do seu carro é: R$ %.2f", valor)
}
