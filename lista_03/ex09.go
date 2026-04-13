package main

import f "fmt"

func main() {
	type aluno struct {
		nota1, nota2, i float32
	}
	var (
		media, mediaturma float32
		i                 int
		n                 aluno
		lista             []aluno

		reprovado, exame, aprovado = 0, 0, 0
	)

	for {
		f.Print("Digite -1 para finalizar\n")
		f.Printf("Insira a 1ª nota do %d° aluno: ", i+1)
		f.Scan(&n.nota1)
		f.Printf("Insira a 2ª nota do %d° aluno: ", i+1)
		f.Scan(&n.nota2)
		if n.nota1 > 10 || n.nota2 > 10 {
			f.Print("NOTA INVÁLIDA, NOTA 0 a 10\n")
			break
		}
		if n.nota1 < 0 {
			break
		}
		lista = append(lista, n)

		i = i + 1

	}

	for i, a := range lista {
		media = (a.nota1 + a.nota2) / 2

		f.Printf("A média do %d aluno foi %.2f\n", i+1, media)

		if media <= 3 {
			f.Print("Reprovado\n\n")
			reprovado = reprovado + 1
		}

		if media > 3 && media <= 7 {
			f.Print("Exame\n\n")
			exame = exame + 1
		}

		if media > 7 && media <= 10 {
			f.Print("Aprovado\n\n")
			aprovado = aprovado + 1
		}
		mediaturma = mediaturma + media
	}

	turma := len(lista)
	mediaturma = mediaturma / float32(turma)

	f.Printf("Total de aprovados: %d\n", aprovado)
	f.Printf("Total de exame: %d\n", exame)
	f.Printf("Total de reprovado: %d\n", reprovado)
	f.Printf("A media da turma foi: %.2f", mediaturma)
}
