package main

import f "fmt"

func main() {
	//CONTROLE BANCÁRIO
	var (
		conta          [10]int
		saldo          [10]float64
		opcao, entrada int
		valor, soma    float64
	)

	busca := make(map[int]float64)

	for i, _ := range conta {
		f.Printf("Conta: ")
		f.Scan(&conta[i])
		f.Printf("Saldo: ")
		f.Scan(&saldo[i])
		busca[conta[i]] = saldo[i] // Todo saldo está atrelada a uma conta
	}

	for {
		for {
			f.Print(
				`		
			MENU
	1 - Efetuar Depósito
	2 - Efetuar Saque
	3 - Consultar o ativo bancário
	4 - Finalizar o Programa
	`)
			f.Scan(&opcao)
			switch opcao {
			case 1:
				f.Print("\tDepósito\n")
				f.Print("Conta: ")
				f.Scan(&entrada)

				_, eValida := busca[entrada] // Retorna se a entrada(conta) está na lista: true or false

				if eValida {
					f.Printf("Valor: ")
					f.Scan(&valor)
					busca[entrada] = busca[entrada] + valor

				} else {
					f.Print("Erro! Conta não encontrada!\n")
					break
				}
			case 2:
				f.Print("\tSaque\n")
				f.Print("Conta: ")
				f.Scan(&entrada)

				_, eValida := busca[entrada] // Retorna se a entrada(conta) está na lista: true or false

				if eValida {
					f.Printf("Valor do Saque: ")
					f.Scan(&valor)

					if (busca[entrada] - valor) >= 0 {
						busca[entrada] = busca[entrada] - valor
					} else {
						f.Print("Saldo insuficiente!")
						break
					}

				} else {
					f.Print("Erro! Conta não encontrada!\n")
					break
				}
			case 3:
				f.Print("\n\tATIVO BANCÁRIO\n")
				soma = 0
				for _, saldo := range busca {
					soma = soma + saldo
				}
				f.Printf("Total ativos: R$ %.2f\n", soma)
			case 4:
				break
			default:
				break

			}
			if opcao == 4 {
				break
			} else {
				continue
			}

		}
		if opcao == 4 {
			break
		} else {
			continue
		}

	}
}
