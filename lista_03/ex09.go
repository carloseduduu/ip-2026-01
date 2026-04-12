package main

import f "fmt"

func main() {
	type aluno struct {
		nota1, nota2, i float32
	}
	var (
		media float32
		i     int
		n     aluno
		lista []aluno
	)

	for {
		f.Print("Digite -1 para finalizar\n")
		f.Printf("Insira a 1ª nota do %d° aluno: ", i+1)
		f.Scan(&n.nota1)
		f.Printf("Insira a 2ª nota do %d° aluno: ", i+1)
		f.Scan(&n.nota2)
		if n.nota1 < 0 {
			break
		}
		lista = append(lista, n)

		i = i + 1

	}

	for i, a := range lista {
		media = (a.nota1 + a.nota2) / 2
		f.Printf("A média do %d aluno foi %.f\n", i+1, media)
	}
}
