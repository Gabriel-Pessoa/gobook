package main

/* O tempo de vida de uma variável é o intervalo de tempo durante o qual ela existe enquanto o programa executa. O tempo de vida de uma variável de nível de pacote
é toda a execução de um programa. Em comparação, as variáveis locais têm tempos de vida dinâmicos: uma nova instância é criada sempre que a instrução de declaração
é executada, e a variável vive até se tornar inacessível, momento em que sua área de armazenagem pode ser reciclada. Parâmetros e resultados de função também são
variáveis locais; eles são criados sempre que a função que os engloba é chamada.
	Como o coletor de lixo sabe que a área de armazenagem de uma variável pode ser reciclada? A ideia básica é que toda variável de nível de pacote e toda variável local
de cada função função ativa no momento podem, possivelmente, ser o início ou a raiz de um caminho (path) para a variável em questão, seguindo ponteiros e outros tipos
de referência que, em última instância, levam à variável. Se esse caminho não existir, é sinal de que a variável se tornou inacessível, portanto não poderá mais afetar o
restante do processamento.
	Como o tempo de vida de uma variável é determinado somente pelo fato de ela ser ou não acessível, uma variável local pode viver mais que uma única iteração do loop que
a engloba. ela pode continuar existindo mesmo após a função que a engloba ter retornado.
	Um compilador pode optar por alocar variáveis locais na heap ou na pilha, mas essa escolha, talvez surpreendentemente, não é determinada pelo fato de var ou
new ter sido usada para declarar a variável.
*/

// Ex:
// for t := 0; t < count; t++ { // a variável t é criada sempre que o loop "for" inicia, e as novas variáveis x e y são criadas a cada iteração do loop
// 	x := logic(t)
// 	y := logic(x)
// }

/* Nesse caso, x (no exemplo 1) deve ser alocada no heap porque continua acessível a partir da variável global depois que f retornar, apesar de ser decladarada como uma
variável local; dizemos que x escapa de f. Por outro lado, quando g retorna, a variável *y torna-se inacessível e pode ser reciclada. Como y não escapa de g, é seguro
para o compilador alocar *y na pilha, apesar de ter sido alocada com new.*/

// Ex 1
// var global *int

// func f() {
// 	var x int
// 	x = 1
// 	global = &x
// }

// Ex 2
// func g() {
// 	y := new(int)
// 	*y = 1
// }
