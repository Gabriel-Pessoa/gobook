/* Outra maneira de criar uma variável é com a função embutida new. A expressão new(T) cria uma variável sem nome do tipo T, inicializa-a com o valor zero de T
e devolve  seu endereço, que é um valor do tipo *T. */
package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)

	// cada chamada a new devolve uma variável distinta com um endereço único:

	t := new(int)
	u := new(int)

	fmt.Println(t == u) // false
}

// Uma variável criada com new não é diferente de uma variável local comum cujo endereço é acessado, exceto pelo
// fato de não haver necessidade de inventar (new declarar) um nome descartável (dummy), e podemos usa new(T) em uma expressão. Desse modo, new é apenas uma.

// func newInt() *int {
// 	return new(int)
// }

// func newInt2() *int {
// 	var dummy int
// 	return &dummy
// }
