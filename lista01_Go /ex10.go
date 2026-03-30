package main

import f "fmt"

func main() {
	var matriz [2][2]int
	//var a, b, c, d int
	f.Scan(&matriz[0][0])
	f.Scan(&matriz[0][1])
	f.Scan(&matriz[1][0])
	f.Scan(&matriz[1][1])
	A := matriz[0][0]
	B := matriz[0][1]
	C := matriz[1][0]
	D := matriz[1][1]
	det := A*D - B*C
	//f.Printf("A = %d\nB = %d\nC = %d\nD = %d\n", A, B, C, D)
	f.Printf("O VALOR DO DETERMINANTE E = %.2f\n", float64(det))
}
