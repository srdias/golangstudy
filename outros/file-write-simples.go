package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("file-write-simples.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString("Linha 1 - AAA\n")
	f.WriteString("Linha 2 - BBB\n")
	f.WriteString("Linha 3 - CCC\n")

	fmt.Println("done")
}
