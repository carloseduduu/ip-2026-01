package main

import f "fmt"

func main() {

	var n1, limite int

	f.Scan(&n1, &limite)
	if n1%2 == 0 {

		for i := 0; i < limite; i++ {
			f.Print(n1, "\n")
			n1 += 2
		}

	} else {

		f.Print("O PRIMEIRO NUMERO NAO E PAR\n")

	}
}
