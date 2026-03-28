package main

import f "fmt"

func main() {
	var n1, n2, n3, media, aluno float64
	var conceito, status string
	f.Print("Insira o número do aluno: ")
	f.Scan(&aluno)
	f.Print("Insira suas notas e média de atividade: ")
	f.Scan(&n1, &n2, &n3, &media)
	mf := (n1 + (n2 * 2) + (n3 * 3) + media) / 7

	if mf >= 9 && mf <= 10 {
		conceito = "A"
		status = "APROVADO"
	}

	if mf >= 7.5 && mf < 9 {
		conceito = "B"
		status = "APROVADO"
	}

	if mf >= 6 && mf < 7.5 {
		conceito = "C"
		status = "APROVADO"
	}

	if mf >= 4 && mf < 6 {
		conceito = "D"
		status = "REPROVADO"
	}

	if mf < 4 {
		conceito = "E"
		status = "REPROVADO"
	}

	f.Printf("Número do Aluno: %.f\n", aluno)
	f.Printf("Suas notas foram:\n%.1f\n%.1f\n%.1f\n", n1, n2, n3)
	f.Printf("A média dos exercícios foi\n%.1f\n", media)
	f.Printf("Seu conceito foi: %s\n", conceito)
	f.Printf("Seu resultado foi: %s\n", status)

}
