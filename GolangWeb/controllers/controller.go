package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	p "../models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

//Index : index
func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", p.BuscaTodosOsProdutos())
}

//New : new
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//Insert : insert
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoToFloat, err := strconv.ParseFloat(preco, 64)
		checkErr(err)
		qntdToInt, err := strconv.Atoi(quantidade)
		checkErr(err)

		p.AdicionarProduto(nome, descricao, precoToFloat, qntdToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
