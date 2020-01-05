package main

import (
	"log"
	"net/http"

	"./routes"
)

func main() {
	routes.CarregaRotas()
	log.Println(http.ListenAndServe(":8000", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
