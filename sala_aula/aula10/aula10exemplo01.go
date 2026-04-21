package main

import (
	"fmt"
	f "fmt"
)

func main() {

	type Criatura struct {
		Nome   string
		idade  int
		altura float64
	}
	var dados Criatura

	f.Scan(&dados.Nome)

	/*c := Criatura{
		Nome:   "Jac and friends",
		idade:  2,
		altura: 1.8,
	}*/

	fmt.Printf("O nome da criatura é %v", dados.Nome)
}
