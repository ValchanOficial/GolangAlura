package contas

import "fmt"

//ContaCorrente : contaCorrente
type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

//Sacar : sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return fmt.Sprintf("[SAQUE] Saque realizado com sucesso. Saldo: %.2f.", c.Saldo)
	}
	return fmt.Sprintf("[SAQUE] Saldo insuficiente. Saldo: %.2f.", c.Saldo)
}

//Depositar : depositar
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return fmt.Sprintf("[DEPÓSITO] Valor depositado. Saldo: %.2f.", c.Saldo)
	}
	return fmt.Sprintf("[DEPÓSITO] Não é possível depositar um valor menor do que zero. Valor: %.2f.", valorDoDeposito)
}

//Transferir : transferir
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		fmt.Printf("[TRANSFERÊNCIA] Valor transferido com sucesso. Saldo: %.2f\n", c.Saldo)
		return true
	}
	fmt.Printf("[TRANSFERÊNCIA] Não é possível transferir. Saldo: %.2f - Transferência: %.2f\n", c.Saldo, valorDaTransferencia)
	return false
}
