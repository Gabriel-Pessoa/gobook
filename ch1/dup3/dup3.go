// As versões antieriores de dup funcioname em um modo "streaming", em que a entrada é lida e separada em linhas
// conforme for necessário, portanto, em princípio, esses programas podem lidar com uma quantidade qualquer de
// dados de entrada. Uma alternativa é ler toda a entrada na memória em um só grande bloco. Nessa versão introduz
// a função readFile (do pacote io/ioutil), que lê todo o conteúdo de um arquivo nomeado, e strings.Split, que separa
// uma string em fatia de substrings.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}

		// ReadFile devolve uma fatia de bytes que deve ser convertida em uma string para que possa ser separada
		// por strings.Split
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
