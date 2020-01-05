package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Conecta : conecta com o Banco De Dados
func Conecta() *sql.DB {
	conn := "user=postgres dbname=alura_loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	log.Println("[Database] - Connected to Database!")
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
