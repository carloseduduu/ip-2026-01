package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var (
		n int
	)

	f.Scan(&n)
	fmt.Println(log2(n))

}

func log2(valor int) int {

	if valor <= 1 {
		return 0
	}

	return 1 + log2(valor/2)
}


/* log de 8 na base 2

Log(2)8 = 1 + (Log(2)4 = 1 + (Log(2)2 = 1 + Log(2)1)
--------------------------------------------------î-------
Aqui o sistema retorna 0
-----------------------------------------------------
Log(2)8 = 1 + [(Log(2)4 + 1 + (Log(2)2 + 1 + 0)]
Log(2)8 = 1 + [(Log(2)4 + 1 + (0 + 1 + 0)]
Log(2)8 = 1 + [0 + 1 + (0 + 1 + 0)]
Log(2)8 = 3

*/