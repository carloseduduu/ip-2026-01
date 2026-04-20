package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var (
		cpf, cpfentrada string
		contador        = 10
		soma, digito    int
	)

	f.Scan(&cpfentrada)

	cpf = cpfentrada[:9] //Seleciona apenas os 9 primeiros dígitos

	for i := 0; i < 9; i++ { //Executa o cálculo para 9 dígitos

		digito := cpf[i] - '0' //converte string para valor

		soma = soma + (int(digito) * contador) // Soma dos valores pela multiplicação de seu índice
		contador--

	}

	resto := soma % 11

	if resto < 2 {
		digito = 0
	} else {
		digito = 11 - resto
	}

	cpf = cpf + strconv.Itoa(digito) //Concatena os 9 dígitos mais o 10º valor
	f.Printf("\nCPF após geração do primeiro dígito: %s.%s.%s-%s\n", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])

	contador = 11
	soma = 0
	for i := 0; i < 10; i++ {

		digito := cpf[i] - '0'

		soma = soma + (int(digito) * contador) // Soma dos
		//f.Printf("Contador: %v Posição: %v, Digito: %v, soma/multi: %v\n", contador, i, digito, soma)
		contador--

	}

	resto = soma % 11

	if resto < 2 {
		digito = 0
	} else {
		digito = 11 - resto
	}

	cpf = cpf + strconv.Itoa(digito) //Concatena os 10 dígitos mais o 11º valor

	f.Printf("\nCPF após geração do segundo dígito: %s.%s.%s-%s\n\n", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])

	if cpfentrada == cpf {
		f.Println("CPF VÁLIDO!\n")
	} else {
		f.Println("CPF INVÁLIDO!\n")
	}
}
