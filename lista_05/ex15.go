package main

import f "fmt"

func main() {
	var (
		v1 [30]int
		vr [30]int
	)

	for i := 0; i < len(v1); i++ {
		f.Scan(&v1[i])
	}

	for i, v := range v1 {
		if i%2 == 0 {
			vr[i] = v * 2
		} else {
			vr[i] = v * 3
		}
	}

	f.Println(vr)
}
