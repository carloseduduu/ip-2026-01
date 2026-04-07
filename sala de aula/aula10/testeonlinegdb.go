package main

import f "fmt"

func main() {
	var soldados, droids int
	f.Print("Insira a quantidade de soldados: ")
	f.Scan(&soldados)
	f.Print("Insira o número de droids: ")
	f.Scan(&droids)

	if soldados >= 1 && soldados <= 1000 && droids >= 1 && droids <= 1000 {
		switch {
		case soldados+droids > 1000:
			f.Println("N")
		default:
			f.Println("S")
		}
	} else {
		f.Println("INVÁLIDO")
	}

}
