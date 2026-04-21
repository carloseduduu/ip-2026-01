package main

import f "fmt"

func main() {
	var n1 int
	f.Scan(&n1)
	switch {
	case n1%3 == 0 || n1%5 == 0:
		f.Printf("O NUMERO E DIVISIVEL\n")
	default:
		f.Print("O NUMERO NAO E DIVISIVEL\n")
	}
}
