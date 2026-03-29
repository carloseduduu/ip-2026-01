package main

import f "fmt"

func main() {
	var far float64
	var qnt int
	var farenait []float64
	f.Scan(&qnt)

	for i := 0; i < qnt; i++ {

		f.Scan(&far)
		farenait = append(farenait, far)
	}

	for i, v := range farenait {
		celsius := (5 * (farenait[i] - 32)) / 9

		f.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", v, celsius)

	}
}
