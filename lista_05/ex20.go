package main

import (
	f "fmt"
)

func main() {
	var (
		dado [20]int
	)

	freq := make(map[int]int)
	for i, _ := range dado {
		//dado[i] = rand.Intn(7-1) + 1 // Gera número aleatório entre 0 e 6 (dado)
		f.Scan(&dado[i])   //Scaneio tudo
		freq[dado[i]] += 1 // Para cada valor repetido eu somo um na própia posição dele
	}

	for chave, quantidade := range freq {
		f.Printf("Valor %d, Freq: %d\n", chave, quantidade)
	}
}
