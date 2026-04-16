package main

import f "fmt"

func main() {
	for i := 0; i < 10; i++ { // linha
		for j := 0; j < 10; j++ { // coluna

			if i < j { // a coluna acima é sempre linha maior que coluna
				f.Printf("(%v, %v) ", i, j)
			}

		}

		f.Print("\n")
	}

}
