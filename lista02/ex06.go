package main

import f "fmt"

func main() {
	var A, B int
	f.Scan(&A, &B)
	if A%B == 0 {
		f.Printf("Valor %d é divisível por %d.\n", A, B)
	} else {
		f.Printf("Valor %d não é divisível por %d.\n", A, B)
	}
}
