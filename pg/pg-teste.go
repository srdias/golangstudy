package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getenv(s string, d string) string {
	v := os.Getenv(s)
	if v != "" {
		return v
	}
	return d
}

func main() {
	fmt.Printf("testing..\n")

	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", "user=postgres password=admin dbname=postgres port=15432 sslmode=disable")

	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return
	}

	defer db.Close()

	fmt.Println("")
	fmt.Println("Conteudo inicial:")
	selectAll(db)

	fmt.Println("")
	fmt.Println("Inserindo alguns dados:")
	insert(db, 1, "Maria")
	insert(db, 2, "Pedro")
	insert(db, 3, "Joaquim")
	insert(db, 4, "Vicenti")
	selectAll(db)

	fmt.Println("")
	fmt.Println("Alterando o nome da pessoa de codigo 2")
	update(db, 2, "Bernardo")
	selectAll(db)

	fmt.Println("")
	fmt.Println("Excluindo a pessoa de codigo 3")
	delete(db, 3)
	selectAll(db)

	fmt.Println("")
	fmt.Println("Apagando todos os dados da tabela")
	deleteAll(db)
	selectAll(db)

	fmt.Println("")
	fmt.Println("Fim!!")

}

func insert(db *sql.DB, i_pessoas int, nome string) {
	sqlStatement := "INSERT INTO golang.pessoas (i_pessoas, nome) VALUES ($1, $2)"
	_, err := db.Exec(sqlStatement, i_pessoas, nome)
	if err != nil {
		panic(err)
	}
}

func update(db *sql.DB, i_pessoas int, nome string) {
	sqlStatement := "UPDATE golang.pessoas SET nome=$2 where i_pessoas = $1"
	_, err := db.Exec(sqlStatement, i_pessoas, nome)
	if err != nil {
		panic(err)
	}
}

func delete(db *sql.DB, i_pessoas int) {
	sqlStatement := "DELETE FROM golang.pessoas where i_pessoas = $1"
	_, err := db.Exec(sqlStatement, i_pessoas)
	if err != nil {
		panic(err)
	}
}

func deleteAll(db *sql.DB) {
	sqlStatement := "DELETE FROM golang.pessoas"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func selectAll(db *sql.DB) error {
	var stmt *sql.Stmt
	var err error

	stmt, err = db.Prepare("select * from golang.pessoas order by i_pessoas")
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return err
	}

	var rows *sql.Rows

	rows, err = stmt.Query()
	if err != nil {
		fmt.Printf("stmt.Query error: %v\n", err)
		return err
	}

	defer stmt.Close()

	count := 0
	for rows.Next() {
		var i_pessoas int
		var nome string

		err = rows.Scan(&i_pessoas, &nome)
		if err != nil {
			fmt.Printf("rows.Scan error: %v\n", err)
			return err
		}

		count++
		fmt.Printf("%02d. %v\n", i_pessoas, nome)
	}

	if count < 1 {
		fmt.Println("Sem linhas para serem exibidas!")
	}

	return nil
}
