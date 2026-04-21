package main

import f "fmt"

func main() {

	var matricula int
	var horas, inss, imposto float64
	const salario = 788
	const extra = 10

	/* Sabe-se:
	• Salário hora-extra = horas-extras * Valor da Hora-Extra;
	• Salário bruto = 3 * Salário Mínimo + Salário hora-extra;
	• Desconto INSS = 12 % do salário bruto, se salário bruto for maior que R$ 1500,00;
	• Desconto do Imposto de Renda = 20 % do Salário Bruto, se o mesmo for maior que R$ 2000,00;
	• Salário líquido = salário bruto – deduções. */
	f.Println("Qual a sua matrícula?")
	f.Scan(&matricula)
	f.Println("Quantas horas extra trabalhadas?")
	f.Scan(&horas)
	hora_extra := horas * extra
	salario_bruto := (3 * salario) + hora_extra

	if salario_bruto > 1500 {

		//inss = 0.12 * salario_bruto
		inss = (12.0 / 100.0) * salario_bruto

	}
	if salario_bruto > 2000 {

		//imposto = 0.20 * salario_bruto
		imposto = salario_bruto * (20.0 / 100.0)

	}
	descontos := imposto + inss
	salario_liquido := salario_bruto - descontos
	f.Printf("O salário líquido é de R$ %.2f.\n", salario_liquido)

}
