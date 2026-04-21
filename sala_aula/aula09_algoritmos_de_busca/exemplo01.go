package main

import f "fmt"

func main() {
	var (
		l = [10]int{22, 5, 15, 24, 67, 45, 1, 76, 21, 11}
		x int
	)

	f.Print("Insira o valor a ser buscado na lista: ")
	f.Scan(&x)

	resultado := buscasequencial(l, x)
	if resultado >= 0 {
		f.Printf("O seu valor está na posição %d\n", resultado)
	} else {
		f.Printf("O valor %d não está no array\n", x)
	}
}

func buscasequencial(l [10]int, x int) int {
	n := l
	for i := 0; i < len(n); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
