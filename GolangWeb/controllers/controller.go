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
	log.Println("[Controller] - Index")
	temp.ExecuteTemplate(w, "Index", p.BuscaTodosOsProdutos())
}

//New : new
func New(w http.ResponseWriter, r *http.Request) {
	log.Println("[Controller] - New")
	temp.ExecuteTemplate(w, "New", nil)
}

//Insert : insert
func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("[Controller] - Insert")
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

//Delete : delete
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	log.Println("[Controller] - Delete - ID: " + idDoProduto)
	p.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

//Edit : edit
func Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("[Controller] - Edit")
	idProduto := r.URL.Query().Get("id")
	produto := p.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

//Update : update
func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("[Controller] - Update")
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idToInt, err := strconv.Atoi(id)
		checkErr(err)
		precoToFloat, err := strconv.ParseFloat(preco, 64)
		checkErr(err)
		qntdToInt, err := strconv.Atoi(quantidade)
		checkErr(err)

		p.AtualizarProduto(idToInt, nome, descricao, precoToFloat, qntdToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
