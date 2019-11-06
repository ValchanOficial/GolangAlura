package main

import (
	"fmt"

	cliente "./clientes"
	conta "./contas"
)

func pagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteValchan := cliente.Titular{Nome: "Valchan", CPF: "123.111.123.120", Profissao: "Desenvolvedora"}
	contaValchan := conta.ContaCorrente{Titular: clienteValchan, NumeroAgencia: 123, NumeroConta: 123456}
	contaValchan.Depositar(3000)
	fmt.Println(contaValchan.ObterSaldo())

	poupancaValchan := conta.ContaPoupanca{Titular: clienteValchan, NumeroAgencia: 123, NumeroConta: 123456, Operacao: 0}
	poupancaValchan.Depositar(3000)

	pagarBoleto(&contaValchan, 60)
	pagarBoleto(&poupancaValchan, 30)
	pagarBoleto(&poupancaValchan, -30)
	fmt.Println(clienteValchan)
	fmt.Println(contaValchan)
	fmt.Println(poupancaValchan)

	// contaValchan := conta.ContaCorrente{"Valchan", 589, 123456, 4125.5}
	// fmt.Println(contaValchan.Saldo)
	// fmt.Println(contaValchan.Sacar(200.00))
	// fmt.Println(contaValchan.Sacar(6000.00))
	// fmt.Println(contaValchan.Depositar(500))
	// fmt.Println(contaValchan.Depositar(-500))

	// contaSilvia := conta.ContaCorrente{"Silvia", 589, 123456, 300}
	// contaGustavo := conta.ContaCorrente{"Gustavo", 589, 123456, 100}

	// fmt.Println(contaSilvia)
	// fmt.Println(contaGustavo)
	// contaSilvia.Transferir(100, &contaGustavo)
	// contaSilvia.Transferir(800, &contaGustavo)
	// fmt.Println(contaSilvia)
	// fmt.Println(contaGustavo)
}
