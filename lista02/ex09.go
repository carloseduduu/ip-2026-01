package main

import f "fmt"

func main() {
	var v1 float64
	f.Print("Insira o Valor da compra: ")
	f.Scan(&v1)

	if v1 < 10.00 {
		v1 += (v1 * 70 / 100)
	} else if v1 >= 10 && v1 < 30 {
		v1 += (v1 * 50 / 100)
	} else if v1 <= 30.00 && v1 < 50.00 {
		v1 += (v1 * 40 / 100)
	} else if v1 >= 50 {
		v1 += (v1 * 30 / 100)
	}

	f.Printf("O valor da venda é de R$ %.2f", v1)
}
