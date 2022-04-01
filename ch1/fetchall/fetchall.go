/*
Quando uma gorrotina tenta receber ou enviar em um canal, ela fica bloqueada até outra gorrotina tentar a operação correspondente de envio ou recebimento;
quando isso acontece, o valor é transferido e ambas as gorrotinas continuam. Nesse exemplo, cada fetch envia um valor (ch <- expressão) pelo canal ch,
e main recebe todos eles (<-ch)
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// um canal é um mecanismo de comunicação que permite que uma gorrotina passe valores de um tipo especificado a outra gorrotina
	ch := make(chan string)

	// para cada argumento de linha de comando, a instrução go inicia uma nova gorrotina que executa fetch assincronamente
	for _, url := range os.Args[1:] {
		// a função  main executa uma gorrotina e a instrução go cria gorrotina adicionais
		go fetch(url, ch) // inicia gorrotina. Uma gorrotina (goroutine), é uma execução concorrente de função.
	}

	// recebe e exibe linha enviada pela função fetch
	for range os.Args[1:] {
		fmt.Println(<-ch) // recebe do canal ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // enviar para o canal
		return
	}
	//io.Copy lê o corpo da resposta e descarta-o, escrevendo no stream de saída ioutil.Discard.
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // defer??
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // enviar para o canal
		return
	}
	secs := time.Since(start).Seconds()
	// À medida que cada resultado chega, fetch envia uma linha de resumo pelo canal ch
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) // enviar para o canal
}
