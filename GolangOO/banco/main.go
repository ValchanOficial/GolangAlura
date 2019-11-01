package main

import (
	"fmt"

	c "./contas"
)

func main() {
	contaValchan := c.ContaCorrente{"Valchan", 589, 123456, 4125.5}
	fmt.Println(contaValchan.Saldo)
	fmt.Println(contaValchan.Sacar(200.00))
	fmt.Println(contaValchan.Sacar(6000.00))
	fmt.Println(contaValchan.Depositar(500))
	fmt.Println(contaValchan.Depositar(-500))

	contaSilvia := c.ContaCorrente{"Silvia", 589, 123456, 300}
	contaGustavo := c.ContaCorrente{"Gustavo", 589, 123456, 100}

	fmt.Println(contaSilvia)
	fmt.Println(contaGustavo)
	contaSilvia.Transferir(100, &contaGustavo)
	contaSilvia.Transferir(800, &contaGustavo)
	fmt.Println(contaSilvia)
	fmt.Println(contaGustavo)
}
