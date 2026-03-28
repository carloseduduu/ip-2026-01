package main

import (
	f "fmt"
	m "math"
)

func main() {
	var n1 float64
	f.Scan(&n1)
	if n1 >= 0 {
		raiz := m.Sqrt(n1)
		f.Printf("%.2f", raiz)
	} else {
		f.Printf("%.2f", m.Pow(n1, 2))
	}

}
