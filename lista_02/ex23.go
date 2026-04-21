package main

import f "fmt"

func main() {
	var idade int
	f.Println("Qual a sua idade?")
	f.Scan(&idade)

	switch {
	case idade < 16:

		f.Println("Não-eleitor")

	case idade >= 18 && idade <= 65:

		f.Println("Eleitor obrigatório")

	case idade >= 16 && idade < 18 || idade > 65:

		f.Println("Eleitor facultativo")

	}

}
