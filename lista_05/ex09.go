package main

import f "fmt"

func main() {
	var (
		altura [10]int
		soma   int
	)

	for i, _ := range altura {
		f.Scan(&altura[i])
		soma += altura[i]
	}

	media := soma / cap(altura)

	for _, v := range altura {
		if v > media {
			f.Printf("Altura acima da média: %d\n", v)
		}
	}

}
