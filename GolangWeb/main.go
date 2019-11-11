package main

import (
	"log"
	"net/http"

	"./routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
