package main

import (
	"fmt"
)

func f() {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // variável local f encobre a função f do nível de pacote
	fmt.Println(g) // variável do nível de pacote
	// fmt.Println(h) // erro de compilação: h não está definido

	x := "hello!"

	// nem todos os blocos léxicos correspondem a sequências de instruções explicitamente delimitadas por chaves;
	// alguns são implícitos. Esse loop for cria dois blocos léxicos: o bloco explícito para o corpo do loop
	// e um bloco implícito que, adicionalmente, engloba as variáveis declaradas pela cláusula de inicialização, como i.
	for i := 0; i < len(x); i++ {
		x := x[i] // refere-se a x do bloco mais externo
		if x != '!' {
			x := x + 'A' - 'a'  // refere-se a x do bloco mais externo
			fmt.Printf("%c", x) // uma letra por iteração
		}
	}

	// Como nos loops for, as instruções if e switch também criam blocos implícitos, além dos blocos de seus corpos. O código da cadeia
	// if-else a seguir mostra o escopo de x e de y
	if x := 1; x == 0 {
		fmt.Println(x)
	} else if y := 2; x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
}

// Declaração curtas de variáveis exigem consciência de escopo.
var cwd string

func init() {

	/* como new cmw nem err estão declaradas no bloco da função init, a instrução := declara ambas como variáveis locais.
	 A declaração interna de cwd deixa a declaração externa inacessível, portanto a instrução não atualiza a variável cwd do nível de pacote como se pretendia
		cwd, err := os.Getwd() // erro de compilação: cwd não é usado
	if err != nil {
		log.Fatal("os.Getwd failed: %v", err)
	}
	*/
}
