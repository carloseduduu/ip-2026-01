package main

import (
	f "fmt"
)

/* RAÍZES DISTINTAS
Condição: delta > 0

RAIZ ÚNICA
Condição: delta = 0

RAÍZES IMAGINÁRIAS
Condição: delta < 0 */

func main() {
	var A, B, C int
	var delta float64
	f.Println("Insira os coeficientes A, B e C")
	f.Scan(&A, &B, &C)

	delta = float64((B * B) - 4*A*C)

	if delta > 0 {
		f.Print("RAÍZES DISTINTAS")
	} else if delta == 0 {
		f.Print("RAÍZ ÚNICA")
	} else {
		f.Print("RAÍZES IMAGINÁRIAS")
	}

}
