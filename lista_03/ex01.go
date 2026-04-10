package main

import f "fmt"

func pot(base, exp int) int {
	resultado := base
	for i := 1; i < exp; i++ {
		resultado = resultado * base
	}
	return resultado
}

func main() {
	var valor, exp int
	f.Print("Insira um valor Inteiro positivo: ")
	f.Scan(&valor)
	f.Print("Insira o expoente: ")
	f.Scan(&exp)
	if valor <= 0 {
		for valor <= 0 {
			f.Print("Insira um valor Inteiro positivo: ")
			f.Scan(&valor)

			if valor > 0 {
				break
			}
		}
	}

	f.Printf("O valor de %v, elevado à %vª potência é igual a: %v\n", valor, exp, pot(valor, exp))

}
