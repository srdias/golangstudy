package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var selectStatement = "select * from golang.pessoas"

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

	err = selectAll(db)
	if err != nil {
		fmt.Printf("select error: %v\n", err)
		return
	}
}

func insert(db *sql.DB) {
	var i_pessoas int
	var nome string
	i_pessoas = 1000
	nome = "João"
	sqlStatement := "INSERT INTO golang.pessoas (i_pessoas, nome) VALUES ($1, $2)"
	_, err := db.Exec(sqlStatement, i_pessoas, nome)
	if err != nil {
		panic(err)
	}
}

func update(db *sql.DB) {
	var i_pessoas int
	var nome string
	i_pessoas = 1000
	nome = "Marcos"

	sqlStatement := "UPDATE golang.pessoas SET nome=$2 where i_pessoas = $1"
	_, err := db.Exec(sqlStatement, i_pessoas, nome)
	if err != nil {
		panic(err)
	}
}

func delete(db *sql.DB) {
	var i_pessoas int
	i_pessoas = 1000
	sqlStatement := "DELETE FROM golang.pessoas where i_pessoas = $1"
	_, err := db.Exec(sqlStatement, i_pessoas)
	if err != nil {
		panic(err)
	}
}

func selectAll(db *sql.DB) error {
	var stmt *sql.Stmt
	var err error

	stmt, err = db.Prepare(selectStatement)
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

	for rows.Next() {
		var i_pessoas int
		var nome string

		err = rows.Scan(&i_pessoas, &nome)
		if err != nil {
			fmt.Printf("rows.Scan error: %v\n", err)
			return err
		}

		fmt.Printf("%02d. %v\n", i_pessoas, nome)
	}

	return nil
}
