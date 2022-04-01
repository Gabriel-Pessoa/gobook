/*
	Esta versão conta o número de requisições;
	uma requisição ao URL /count devolve a contagem até agora, excluindo as próprias requisições /count. Internamente, o servidor executa
	o handler para cada requisição de entrada de uma gorrotina separada para que possa servir múltiplas requisições simultaneamente.
	Entretanto, se duas requisções concorrentes tentarem atualizar count ao mesmo tempo, pode ser que ele não seja incrementado consistemente;
	o pragrama teria um bug sério chamado race condition. Para evitar esse problema, devemos garantir que, no máximo, uma gorrotina acesse a
	variável em determinado instante, que é o propósito das chamadas a mu.Lock() e a mu.Unlock().
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// handler ecoa componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter ecoa o número de chamadas até agora
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
