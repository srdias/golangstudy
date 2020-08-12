package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	b := []byte("Linha 1\nLinha 2\nLinha 3\n")

	// write the whole body at once
	err := ioutil.WriteFile("file-write-simples-2.txt", b, 0644)
	if err != nil {
		panic(err)
	}

	// read the whole file at once
	b, err = ioutil.ReadFile("file-write-simples.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Conteudo lido:\n[%s]", b)

}
