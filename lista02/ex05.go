package main

import f "fmt"

func main() {
	var n1 int

	f.Scan(&n1)
	if n1%5 == 0 {
		f.Println("É divisível por 5")
	} else {
		f.Println("Não é divisível por 5")
	}
}
