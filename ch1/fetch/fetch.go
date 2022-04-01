// O mínimo necessário para recperar informações via HTTP, a seguir apresentamos um programa
// simples, insopirado no valioso utilitário curl.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		res, err := http.Get(url) // faz uma requisição HTTP e, se não houver erro, devolve o resultado na estrutura de resposta resp.
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1) // qualquer que seja o caso de erro, os.Exit(1) faz o processo sair com um código de status igual a 1
		}

		// res.Body contém a resposta do servidor na forma de um strem pronto para leitura
		b, err := ioutil.ReadAll(res.Body) // ioutil.ReadAll lê toda a resposta e o resultado é armazenado em b
		res.Body.Close()                   // o stream body é fechado para evitar vazamentos de recursos
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
