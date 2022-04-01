package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countsLines(os.Stdin, counts)
	} else {
		for _, arg := range files {

			// os.Open devolve dois valores: um, é um arquivo aberto (*os.File) usado em leituras subsequentes por Scanner.
			// o segundo, é um valor embutido error. Se error for igual ao valor embutido nil, o arquivo foi aberto com sucesso.
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// quando o mapa (map) é passado para uma função, ela recebe uma cópia da referência, portanto qualquer alteração que a função chamada
// fizer na estrutura de dados subjacente também será vísivel pela referência ao mapa por quem fez a chamada. Nesse exemplo os valores
// inseridos no mapa counts por countLines são visto por main.
func countsLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Nota: ignorando erros em potencial de input.Err()
}
