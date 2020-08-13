package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	conteudoGravarByteArray := []byte("Linha 1\nLinha 2\nLinha 3\n")

	// write the whole body at once
	errGravando := ioutil.WriteFile("file-write-simples.txt", conteudoGravarByteArray, 0644)
	if errGravando != nil {
		panic(errGravando)
	}

	// read the whole file at once
	conteudoLidoByteArray, errLeitura := ioutil.ReadFile("file-write-simples.txt")
	if errLeitura != nil {
		panic(errLeitura)
	}

	conteudoLidoString := string(conteudoLidoByteArray)
	fmt.Printf("Conteudo lido:\n[%s]", conteudoLidoString)

}
