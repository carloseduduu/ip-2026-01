package main

import f "fmt"

func main() {

	type empregados struct {
		ID    int
		meses int
	}
	var valor empregados
	lista := make([]empregados, 0, 100)

	for i := 0; i < cap(lista); i++ {
		f.Scan(&valor.ID, &valor.meses)

		if valor.ID == 0 && valor.meses == 0 {
			break
		}

		lista = append(lista, valor)
	}

	n := len(lista)

	for i := 0; i < n-1; i++ {

		/*Se você tem 5 funcionários, após organizar 4 deles
		nos lugares certos, o 5º obrigatoriamente estará na posição correta.
		Por isso, não precisamos de uma 5ª rodada.*/

		for j := 0; j < n-i-1; j++ {

			if lista[j].meses > lista[j+1].meses {
				temp := lista[j+1]
				lista[j+1] = lista[j]
				lista[j] = temp
			}
		}
	}

	recentes := 3

	if len(lista) < recentes {
		recentes = len(lista)
	}

	f.Printf("Os %d empregados mais recentes são:\n", recentes)
	for i := 0; i <= recentes-1; i++ {
		f.Printf("ID: %d \t Meses: %d\n", lista[i].ID, lista[i].meses)
	}
}
