package main

// sintaxe
// var nome tipo = expressão
var idade int = 26

// pode-se omitir a expressão ou valor
var idade2 = 26
var idade3 int // armazenará o valor zero que garante que uma variável será inicializada

// declarar e, opcionalmente, inicializar um conjunto de variáveis em uma única declaração.
var i, j, k int // int, int, int

var b, f, s = true, 2.3, "four" // bool, float64, string

// Os inicalizadores podem ser valores literais ou expressões arbitrárias.

// Short variable declaration, por serem curtas e flexíveis, elas são usadas para declarar e inicializar a maioria das variáveis locais.
// exs:

// t := 0.0

// f, err := os.Open(x)

// f, err := os.Create(x) // o operador := atuará apenas como atribuição para a variável erro.
