// Echo4 exibe seus argumentos de linhas de comando
package main

import (
	"flag"
	"fmt"
	"strings"
)

// ponteiros são fundamentais para o pacote flag, que usa os argumentos de linha de comando de um programa para definir os valores
// de determinadas variáveis.

// cria uma nova variável flag do tipo bool. Ela aceita três argumentos: o nome da flag ("n"), o valor default da variável (false) e
// uma mensagem que será exibida se o usuário fornecer um argumento inválido, ex: -h ou -help
var n = flag.Bool("n", false, "omit trailing newline") // omite a quebra de linha final que normalmente seria exibida

// de modo semelhante, flag.String aceita um nome, um valor default e uma mensagem, e cria uma variável do tipo string.
var sep = flag.String("s", "", "separator") // separa os argumentos de saídam usando o conteúdo da string em vez de utilizar o caractere de espaço default

// as variáveis sep e n são ponteiros para as variáveis de flag, que devem ser acessadas indiretamente como *n e *sep
func main() {
	// quando o programa é executado, ele deve chamar flag.Parse antes das flag serem usadas para atualizar as variáveis de flag em relação a seus
	// valores default
	flag.Parse() // se flag.Parse() encontrar um erro, ela exibe uma mensagem de uso e chama os.Exit(2) para terminar o programa.

	// os argumentos que não são flag estão disponíveis em flag.Args() como uma fatia de strings
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
