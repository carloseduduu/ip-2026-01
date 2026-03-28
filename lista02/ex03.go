package main

import f "fmt"

func main() {
	var v1, v2 int
	f.Scan(&v1, &v2)
	soma := v1 + v2
	if soma > 20 {
		f.Printf("%d", soma+8)
	} else {
		f.Printf("%d", soma-5)
	}
}
