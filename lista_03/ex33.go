package main

import f "fmt"

func main() {
	var (
		n1, n2, resto, quociente int
	)
	f.Print("Seja n1/n2, insira separado por expaço o valor de n1 e n2: ")
	f.Scan(&n1, &n2)

	resto = n1 - n2

	for {
		if resto >= n2 {
			resto = resto - n2
			quociente++
		} else if resto == 0 {
			quociente++
			break
		}

	}

	f.Printf("Quociente: %d, Resto: %d\n", quociente, resto)
}
