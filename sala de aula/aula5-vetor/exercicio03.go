package main

import f "fmt"

func main() {
	const limite int = 3
	var numeros [limite]float64
	for i := 0; i < limite; i++ {
		f.Scan(&numeros[i])
	}
	for i := len(numeros); i > 0; i-- {
		f.Print(numeros[i-1])
	}
}
