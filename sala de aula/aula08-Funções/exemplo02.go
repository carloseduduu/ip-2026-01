package main

import f "fmt"

func main() {
	var n1, n2, n3 int
	f.Scan(&n1, &n2, &n3)
	maior := maior(n1, n2, n3)
	f.Printf("O maior valor é %d\n", maior)

}

func maior(valor1, valor2, valor3 int) int {
	if valor1 > valor2 {
		temp := valor1
		valor1 = valor2
		valor2 = temp
	}
	if valor2 > valor1 {
		temp := valor2
		valor1 = valor2
		valor2 = temp
	}
	if valor1 > valor3 {
		temp := valor1
		valor1 = valor3
		valor3 = temp
	}
	return valor3
}
