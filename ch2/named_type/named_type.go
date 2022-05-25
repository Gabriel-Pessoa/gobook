/*
	Um declaração "type" define um novo tipo nomeado que tem o mesmo tipo subjacente de um tipo existente. O tipo nomeado oferece uma maneira de distinguir
	usos diferentes, e talvez imcompatíveis, do tipo subjacente, para que eles não sejam confundidos involuntariamente.

	type nome tipo-subjacente

	As declarações de tipo aparecem com mais frequência no nível de pacote, em que tipo nomeado é visível em todo o pacote e, se for exportado
	(começa com letra maiúscula), o nome será acessível também a outros pacotes.
*/
package main

import "fmt"

// Realiza cálculos com temperaturas em Celsius e em fahrenheit

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// A declaração a seguir é um método de tipo, em que parâmetro c do tipo Celsius aparece antes do nome da função, associa ao tipo Celsius um método chamado String,
// que devolve o valor numérico de c seguido de ºC
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100ºC

	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC)) // 180ºF
	// fmt.Printf("%g\n", boilingF-FreezingC) // erro de compilação: incompatibilidade de tipos

	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true

	//fmt.Println(c == f) // erro de compilação: incompatibilidade de tipos
	fmt.Println(c == Celsius(f)) // true
	// Nesse último caso, apesar de seu nome, a conversão de tipo Celsius(f) não altera o valor de seu argumento, apenas seu tipo.
	// O teste é verdadeiro porque c e f são iguais a zero.

	// método de tipo
	c = FtoC(212.0)

	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // não há necessidade de chamar String
	fmt.Printf("%s\n", c)
	fmt.Println(c)

	fmt.Printf("%g\n", c)   // não chama String
	fmt.Println(float64(c)) // não chama String
}
