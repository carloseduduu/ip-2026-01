package main

import f "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			f.Print("      ")
			if i == j {
				f.Printf("(%v, %v) ", i, j)
			}
		}
		f.Print("\n")
	}
}
