package main

import "fmt"

/* Tipos nomeados: uma declaração type dar um nome a um tipo existente. Como tipos referentes a estruturas muitas vezes são longos, quase sempre
eles são nomeados . Um exemplo familiar é a definição de um tipo Point para um sistema gráfico 2D*/
type Point struct {
	X, Y int
}

func main() {

	var point Point
	fmt.Println(point)
	fmt.Println(point.X)
	fmt.Println(point.Y)

	/*Ponteiros: Go oferece ponteiros, isto é, valores que contêm o endereço de uma variável. Em algumas linguagens, notadamente em C, ponteiros são relativamente irrestritos.
	Em outras outras linguagens, ponteiros são disfarçados como "referências" e não há muito que se possa fazer com eles a não ser passá-los de um lado para o outro.
	Go assume uma posição, de certo modo, intermediária. Os ponteiros são explicitamente visíveis. O operador "&" fornece o endereço de uma variável, e o operador "*"
	recupera a variável à qual o ponteiro se refere, mas não há aritmética com ponteiros. */
	i, j := 42, 2701

	p := &i         // ponteiro para i
	fmt.Println(*p) // lê o valor de  i através do ponteiro
	*p = 21         // seta valor de i através do ponteiro

	p = &j
	*p = *p / 37   // ponteiro para i
	fmt.Println(j) // vê o novo valor de j

	/* Métodos e interfaces: um método é uma função associada a um tipo nomeado; Go é incomum no sentido em que métodos podem ser associados a quase
	todo tipo nomeado. As interfaces são tipos abstratos que nos permitam tratar tipos concretos diferentes da mesma maneira, com base nos métodos
	que eles têm, e não no modo como são representados ou implementados. */

	/* Pacotes: Go vem com uma biblioteca-padrão extensa de pacotes úteis, e a comunidade Go vem criando e compartilhando muitos outros. Programação
	geralmente tem mais a ver com o uso de pacotes existentes que a escrita de um código original por conta própria.*/

	/* Comentários: Já mencionamos comentários para documentação no início de um programa ou de um pacote. Também é considerado um bom estilo escrever um
	comentário antes da declaração de cada função para especificar o seu comportamento. Essas convenções são importantes porque elas são usadas por ferramentas
	como go doc e godoc para localizar e exibir documentação */
}
