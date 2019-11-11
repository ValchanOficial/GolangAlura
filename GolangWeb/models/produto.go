package models

import (
	"log"

	database "../db"
)

//Produto : produto
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//BuscaTodosOsProdutos : buscaTodosOsProdutos
func BuscaTodosOsProdutos() []Produto {
	db := database.Conecta()
	defer db.Close()

	selectProdutos, err := db.Query("SELECT * FROM produtos")
	checkErr(err)

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		checkErr(err)

		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}
	return produtos
}

//AdicionarProduto : adicionarProduto
func AdicionarProduto(nome, descricao string, preco float64, quantidade int) {
	db := database.Conecta()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	checkErr(err)
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
