package main

import f "fmt"

func main() {
	var (
		qnt           int //qunatidade de perguntas e para cada pergunta rodar a busca
		num           int // número de artefatos que vai ser inserido o código
		artefatos_cod []int
		lista_chave   []int
		encontrado    bool
	)

	f.Scan(&num)
	f.Scan(&qnt)

	for i := 0; i < num; i++ {
		valor := 0
		f.Scan(&valor)
		artefatos_cod = append(artefatos_cod, valor)
	}
	f.Println("2ª PARTE")
	for i := 0; i < qnt; i++ {
		lista2 := 0
		f.Scan(&lista2)
		lista_chave = append(lista_chave, lista2)
	}

	for i := 0; i < len(lista_chave); i++ {
		encontrado = false
		for j := 0; j < len(artefatos_cod); j++ {
			if lista_chave[i] == artefatos_cod[j] {
				encontrado = true
				break
			}
		}
		if encontrado {
			f.Println("Yes")
		} else {
			f.Println("No")
		}
	}
}
