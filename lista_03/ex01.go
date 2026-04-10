package main
import f "fmt"

func pot(base, exp int) int {
	for i := 0; i < exp; i++{
		resultado = base
		resultado = resultado*base
	}
	return resultado
}

func main(){
	var valor int
	f.Print("Insira um valor Inteiro positivo: ")
	f.Scan(&valor, &exp)
	if valor <= 0 {
		for valor <= 0 {
		f.Print("Insira um valor Inteiro positivo: ")
		f.Scan(&valor)

		if valor > 0 {break}
		}
	}

	f.Printf("O valor de %v, elevado à %vª potência é igual a: %v", valor, exp, pot(valor,exp))
	
}

