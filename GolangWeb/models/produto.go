package models

import (
	"log"
	"strconv"

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
	log.Println("[Model] - BuscaTodosOsProdutos")
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

		p.ID = id
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
	log.Println("[Model] - AdicionarProduto")
	db := database.Conecta()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	checkErr(err)
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
}

//DeletaProduto : deletaProduto
func DeletaProduto(id string) {
	log.Println("[Model] - DeletaProduto")
	db := database.Conecta()
	defer db.Close()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	checkErr(err)
	deletarProduto.Exec(id)
}

//EditaProduto : editaProduto
func EditaProduto(id string) Produto {
	log.Println("[Model] - EditaProduto")
	db := database.Conecta()
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	checkErr(err)

	p := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		checkErr(err)

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco
	}
	return p
}

//AtualizarProduto : atualizarProduto
func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	log.Println("[Model] - AtualizarProduto - ID: " + strconv.Itoa(id))
	db := database.Conecta()
	defer db.Close()

	atualizaDados, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	checkErr(err)
	atualizaDados.Exec(nome, descricao, preco, quantidade, id)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
