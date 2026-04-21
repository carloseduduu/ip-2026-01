package main

import f "fmt"

func main() {
	var d1, ida int

	f.Print(
		`Qual o seu destino?

1 - Região Norte
2 - Região Nordeste
3 - Região Centro-Oeste
4 - Região Sul

Insira o valor correspondente: `)
	f.Scan(&d1)

	f.Println(
		`Sua viagem é de qual categoria?
1 - Ida
2 - Ida e Volta`)
	f.Scan(&ida)

	switch d1 {
	case 1:
		if ida == 1 {
			f.Println("O valor da passagem de Ida para o destino Região Norte é de R$ 500,00.")
		} else {
			f.Println("O valor da passagem de Ida e Volta para o destino Região Norte é de R$ 900,00.")
		}

	case 2:
		if ida == 1 {
			f.Println("O valor da passagem de Ida para o destino Região Nordeste é de R$ 350,00.")
		} else {
			f.Println("O valor da passagem de Ida e Volta para o destino Região Nordeste é de R$ 650,00.")
		}

	case 3:
		if ida == 1 {
			f.Println("O valor da passagem de Ida para o destino Região Centro-Oeste é de R$ 350,00.")
		} else {
			f.Println("O valor da passagem de Ida e Volta para o destino Região Centro-Oeste é de R$ 600,00.")
		}

	case 4:
		if ida == 1 {
			f.Println("O valor da passagem de Ida para o destino Região Sul é de R$ 300,00.")
		} else {
			f.Println("O valor da passagem de Ida e Volta para o destino Região Sul é de R$ 550,00.")
		}
	}
}
