package main

import f "fmt"

func main() {

	f.Print("\n")
	for linha := 0; linha < 10; linha++ {

		for colunas := 0; colunas < 10; colunas++ {
			f.Printf("(%d,%d) ", linha, colunas)
		}

		f.Print("\n")
	}
	f.Print("\n")
}
