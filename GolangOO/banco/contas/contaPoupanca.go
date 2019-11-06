package contas

import (
	"fmt"

	c "../clientes"
)

//ContaPoupanca : contaPoupanca
type ContaPoupanca struct {
	Titular       c.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

//Sacar : sacar
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("[SAQUE] Saque realizado com sucesso. Saldo: %.2f.", c.saldo)
	}
	return fmt.Sprintf("[SAQUE] Saldo insuficiente. Saldo: %.2f.", c.saldo)
}

//Depositar : depositar
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return fmt.Sprintf("[DEPÓSITO] Valor depositado. Saldo: %.2f.", c.saldo)
	}
	return fmt.Sprintf("[DEPÓSITO] Não é possível depositar um valor menor do que zero. Valor: %.2f.", valorDoDeposito)
}

//ObterSaldo : obterSaldo
func (c *ContaPoupanca) ObterSaldo() string {
	return fmt.Sprintf("[SALDO] Saldo: %.2f.", c.saldo)
}
