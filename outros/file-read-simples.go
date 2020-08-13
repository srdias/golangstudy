package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("world192.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("\nImprimindo somente as 20 primeiras linhas do arquivo e contando quantas linhas existem no arquivo.\n")

	contador := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contador++
		if contador <= 20 {
			fmt.Printf("Linha: %02d. [%v]\n", contador, scanner.Text())
		}
	}

	fmt.Printf("\nTotal de linha do arquivo: %d.\n\n", contador)

}
