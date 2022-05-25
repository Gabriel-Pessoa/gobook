package main

import "fmt"

const boilingF = 212.0 // é uma declaração de nível de pacote ( assim como a função main).

func main() {
	// as variáveis f e c são locais à função main
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
}
