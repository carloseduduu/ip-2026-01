package main

import f "fmt"

func main() {
	var n1, razao, limite, soma int
	f.Scan(&n1, &razao, &limite)
	soma = 0
	for i := 1; i <= limite; i++ {
		soma = soma + n1
		n1 = n1 + razao

	}
	f.Println(soma)
}
