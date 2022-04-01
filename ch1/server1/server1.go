// Para rodar esse código: go run server1.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // cada requisição chama handler
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// handler ecoa o comportamento Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
