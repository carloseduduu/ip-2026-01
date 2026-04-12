package main

import f "fmt"

func main() {
	var (
		valor, qnt, qntimpar, somapares   int
		lista_valores                     []float64
		soma, menor, maior, percent_impar float64
		i                                 = 0
		media, mediapar                   float64
	)
	// entrada dos valores
	for i = 0; valor != 30000; i++ {
		soma = soma + float64(valor)
		f.Print("Insira os valores: ")
		f.Scan(&valor)
		//Encerra ao digitar 30000
		if valor == 30000 {
			break
		}

		lista_valores = append(lista_valores, float64(valor))
	}
	media = soma / float64(i)
	//retorna o maior e o menor valor
	for _, v := range lista_valores {
		maior = lista_valores[0]

		if v > maior {
			maior = v
		}

		menor = lista_valores[0]

		if v < menor {
			menor = v
		}
	}

	for _, v := range lista_valores {
		if int(v)%2 == 0 {
			somapares += int(v)
			qnt += 1
		} else {
			qntimpar += 1
		}
	}
	mediapar = float64(somapares) / float64(qnt)
	total := len(lista_valores)
	percent_impar = (float64(qntimpar) / float64(total)) * 100

	f.Printf("A soma é %.2f\nA quantidade é %d\nA média é %.2f\nO menor é %.f\nO maior é %.f\n", soma, i, media, menor, maior)
	if mediapar > 0 {
		f.Printf("A média dos pares é %.2f\n", mediapar)
	} else {
		f.Print("Não há valores pares\n")
	}
	f.Printf("Percentagem ímpar %.1f%%", percent_impar)
}
