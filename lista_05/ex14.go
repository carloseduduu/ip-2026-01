package main

import f "fmt"

func main() {
	var (
		v1     [10]int
		v2     [10]int
		p1, p2 int
	)
	vr := make([]int, 0, len(v1)+len(v2))
	for i := 0; i < cap(v1); i++ {
		f.Scan(&v1[i])
	}

	for i := 0; i < cap(v2); i++ {
		f.Scan(&v2[i])
	}

	for i := 0; i < len(v1)+len(v2); i++ {

		if i%2 == 0 {
			vr = append(vr, v1[p1])
			p1++
		} else {
			vr = append(vr, v2[p2])
			p2++
		}

	}

	f.Println(vr)
}
