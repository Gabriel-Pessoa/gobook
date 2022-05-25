/* Os pacotes em Go têm o mesmo propósito das bibliotecas oum módulos em outras linguagens; eles suportam modularidade, encapsulamento,
compilação separada e reutilização. O código-fonte de um pacote está em um ou mais arquios .go, normalmente em um diretório cujo nome termina
com o path de importação: por exemplo, os arquivos do pacote github.com/Gabriel-Pessoa/gobook/ch1/helloworld são armazenados no diretório
$GOPATH/src/github.com/Gabriel-Pessoa/gobook/ch1/helloworld.
	Cada pacote serve como um name space (espaço de nomes) separado para suas declarações. Na pacote image, por exemplo, o identificador Decode
refere-se a uma função diferente do identificador de mesmo nome no pacote unicode/utf16. Para referenciar uma função fora de seu pacote, devemos
qualificar o identificador para deixar mais explícito se queremos dizer image.Decode ou utf16.Decode
	Em Go, uma regra simples determina quais identificadores são exportados e quais não são: identificadores exportados começam com letra maiúscula.
*/
package main

import (
	"fmt"

	"github.com/Gabriel-Pessoa/gobook/ch2/package-files/tempconv"
)

func main() {
	// Acessando constante pública de um pacote importado
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)

	// Convertendo uma temperatura em Celsius para Fahrenheit
	fmt.Println(tempconv.CtoF(tempconv.BoilingC))
}
