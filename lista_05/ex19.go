package main

import f "fmt"

func main() {
	var (
		num [10]int
		div [5]int
	)

	for i, _ := range num {
		f.Scan(&num[i])
	}

	for i, _ := range div {
		f.Scan(&div[i])
	}

	for i := 0; i < len(num); i++ {
		f.Printf("Número %d:\n", num[i])
		for j := 0; j < len(div); j++ {

			if num[i]%div[j] == 0 {

				f.Printf("\tDívisível por %d na posição %d\n", div[j], j)

			}
		}
	}

}
