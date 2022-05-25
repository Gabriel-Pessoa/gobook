/*
	Uma variável é um item de armazenagem que contém um valor. Variáveis são identificadas por um nome, por exemplo, x,
	mas muitas variáveis são identificadas somente por expressões.

	Um ponteiro é o endereço de uma variável. Desse modo, um ponteiro é o local em que um valor é armazenado. Nem todo valor
	tem um endereço, mas toda a variável tem. Com um ponteiro, podemos ler ou atualizar o valor de uma variável indiretamente,
	sem usar ou sequer saber o nome da variável, se é que ela tem nome.
*/

package main

import "fmt"

func main() {
	var x int
	var p *int // declaração; ou poderia p := &x. Evitaria a linha seguinte de atribuição

	p = &x // p, do tipo *int, aponta para x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	// Variáveis podem ser descritas como valores endereçáveis. Expressões que representam variáveis são as únicas expressões
	// Expressões que representam variáveis são as únicas expressões para as quais o operador "&"(endereço-de) é aplicável.

	// O valor zero de um ponteiro de qualquer tipo é nil. O teste p != nil é verdadeiro se p aponta para uma variável.
	// Ponteiros são comparáveis; dois ponteiros serão iguais se e somente se apontarem para a mesma variável ou ambos forem nil.

	var a, b int

	fmt.Println(&a == &a, &a == &b, &a == nil) // true, false, false

	// é perfeitamente seguro para uma função devolver o endereço de uma variável local, criada em particular a f, existirá mesmo
	// depois de a chamada ter retornado, e o ponteiro point continuará fazendo referência a ela
	// var point = f()

	// cada chamada de f devolve um valor distinto
	fmt.Println(f() == f()) // false

	num := 1
	incr(&num)              // efeito: num agora é 2
	fmt.Println(incr(&num)) // efeito: num agora é 3
}

func f() *int {
	v := 1
	return &v
}

// como o ponteiro contém o endereço da variável, passar um argumento do tipo ponteiro a uma função possibilita que a função atualize
// a variável apontada pelo seu argumento e devolve o novo valor da variável de modo que ele pode ser usado em uma expressão
func incr(p *int) int {
	*p++ // incrementa o valor para o qual p aponta; não altera p
	return *p
}

// sempre que usamos o endereço de uma variável ou copiamos um ponteiro criamos novos alias (apelidos), ou maneiras de identificar a mesma variável.
// Na função incr, p é uma alias para o endereço da variável que será passada como argumento da função. É conveniente porque nos permite acessar uma variável
// sem utilizar o seu nome, mas é uma faca de dois gumes: para encontrar todas as instruções que acessam uma variável precisamos conhecer todos os seus aliases.
