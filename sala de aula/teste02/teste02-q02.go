package main

import f "fmt"

func main() {
	var (
		n int
	)
	f.Scan(&n)

	if n >= 1 && n <= 10000000 {

		for n != 1 {
			f.Printf("%d ", n)
			if n%2 == 0 {
				n = n / 2
			} else {
				n = (n * 3) + 1
			}

		}
	}
}
