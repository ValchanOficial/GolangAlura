package contas

import (
	"fmt"

	c "../clientes"
)

//ContaCorrente : contaCorrente
type ContaCorrente struct {
	Titular       c.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

//Sacar : sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("[SAQUE] Saque realizado com sucesso. Saldo: %.2f.", c.saldo)
	}
	return fmt.Sprintf("[SAQUE] Saldo insuficiente. Saldo: %.2f.", c.saldo)
}

//Depositar : depositar
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return fmt.Sprintf("[DEPÓSITO] Valor depositado. Saldo: %.2f.", c.saldo)
	}
	return fmt.Sprintf("[DEPÓSITO] Não é possível depositar um valor menor do que zero. Valor: %.2f.", valorDoDeposito)
}

//Transferir : transferir
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		fmt.Printf("[TRANSFERÊNCIA] Valor transferido com sucesso. Saldo: %.2f\n", c.saldo)
		return true
	}
	fmt.Printf("[TRANSFERÊNCIA] Não é possível transferir. Saldo: %.2f - Transferência: %.2f\n", c.saldo, valorDaTransferencia)
	return false
}

//ObterSaldo : obterSaldo
func (c *ContaCorrente) ObterSaldo() string {
	return fmt.Sprintf("[SALDO] Saldo: %.2f.", c.saldo)
}
