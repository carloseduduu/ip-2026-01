package main

import "fmt"

func main() {
	v := 0
	incrementa(&v)
	fmt.Println(v) // 1
}
func incrementa(x *int) {
	*x++
}
