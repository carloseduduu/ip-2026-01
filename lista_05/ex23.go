package main

import f "fmt"

func main() {
	var (
		janela   = [24]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
		corredor [24]bool
		opcao    int
	)

	f.Print("Qual poltrona deseja?\n")
	f.Print("1 - Janela\n")
	f.Print("2 - Corredor\n")
	f.Scan(&opcao)

	vazias := make([]int, 0, 0) //Var que guarda as poltronas vazias

	switch opcao {
	case 1:
		for i, v := range janela {
			if !v {
				vazias = append(vazias, i)
			}
		}
		if len(vazias) == 0 {
			f.Printf("Não há poltronas livres!")
			break
		}
		f.Printf("Poltronas janela disponíveis\n")
		for _, poltrona := range vazias {
			f.Printf("Poltrona Nº %d\n", poltrona)
		}

	case 2:
		for poltrona, v := range corredor {
			if !v {
				vazias = append(vazias, poltrona)
			}
		}
		if len(vazias) == 0 {
			f.Printf("Não há poltronas livres!")
			break
		}
		f.Printf("Poltronas corredor disponíveis\n")
		for _, poltrona := range vazias {
			f.Printf("Poltrona Nº %d\n", poltrona)
		}
	}
}
