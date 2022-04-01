// O mínimo necessário para recperar informações via HTTP, a seguir apresentamos um programa
// simples, insopirado no valioso utilitário curl.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	const prefix = "http://"

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v code: %v\n", err, res.StatusCode)
			os.Exit(1)
		}

		dst := os.Stdout

		if _, err = io.Copy(dst, res.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		res.Body.Close()

		fmt.Printf("%s", dst)
	}
}
