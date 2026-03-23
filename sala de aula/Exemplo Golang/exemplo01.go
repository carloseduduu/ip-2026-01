package main

import f "fmt"

func main() {
	var l1, l2, l3 float32

	f.Printf("Digite o valor do lado a: \n")
	f.Scan(&l1)
	f.Printf("Digite o valor do lado b: \n")
	f.Scan(&l2)
	f.Printf("Digite o valor do lado c: \n")
	f.Scan(&l3)

	//Validação dos valores
	if l1 >= (l2+l3) || l2 >= (l1+l3) || l3 >= (l1+l2) {
		f.Print("VALORES INVÁLIDOS")
	} else if l1 == l2 && l2 == l3 {
		f.Print("É um triângulo equilátero")
	} else if l1 != l2 && l2 != l3 {
		f.Print("É um triângulo Escaleno")
	} else if l1 == l2 || l1 == l3 {
		f.Print("É um triângulo isóceles")
	}
}
