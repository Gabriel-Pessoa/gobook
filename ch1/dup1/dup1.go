package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) // cria variável que se refere a um bufio.Scanner

	// o scanner lê da entrada-padrão do programa.
	// o input.scan lê a próxima linha e remove o caractere de quebra de linha do final;
	// o resultado pode ser obtido por meio de chamada a input.Text().
	// a função Scan devolve true se houver uma linha, e false quando não houver mais dados de entrada
	for input.Scan() {
		counts[input.Text()]++
	}

	// Ignorando erros em potencial do input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
