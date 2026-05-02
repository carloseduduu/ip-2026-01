package main

import f "fmt"

func main() {
	var (
		n     int
		A     []int
		m     int
		B     []int
		l     int
		C     []int
		q     int
		X     []int
		valor int
	)

	f.Scan(&n)
	for i := 0; i < n; i++ {
		f.Scan(&valor)
		A = append(A, valor)
	}

	f.Scan(&m)
	for i := 0; i < m; i++ {
		f.Scan(&valor)
		B = append(B, valor)
	}

	f.Scan(&l)
	for i := 0; i < l; i++ {
		f.Scan(&valor)
		C = append(C, valor)
	}

	f.Scan(&q)
	for i := 0; i < q; i++ {
		f.Scan(&valor)
		X = append(X, valor)
	}

	for i := 0; i < len(X); i++ {
		
		for _, valor = range A {
			
		}
	}

}
