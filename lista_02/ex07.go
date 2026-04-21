package main

import f "fmt"

func main() {
	var A, B, C int
	f.Scan(&A, &B, &C)

	if A > C {
		temp := C
		C = A
		A = temp
	}
	if A > B {
		temp := A
		A = B
		B = temp
	}
	if B > C {
		temp := B
		B = C
		C = temp
	}

	f.Println(A, B, C)
}
