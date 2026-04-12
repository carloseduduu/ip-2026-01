package main

import f "fmt"

func main() {
	var (
		soma int
		
	)

	for i := 0; i <= 20; i++ {
		soma += i
		f.Println(i)
	}

	f.Printf("A soma de todos é %d", soma)
}
