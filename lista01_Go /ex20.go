package main

import f "fmt"

func main() {
	var hora, min, seg int
	f.Scan(&hora, &min, &seg)
	seg = (hora * 3600) + (min * 60) + seg
	f.Printf("O TEMPO EM SEGUNDO E = %d\n", seg)
}
