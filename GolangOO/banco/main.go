package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *contaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

// função variádica
func somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	contaValchan := contaCorrente{"Valchan", 589, 123456, 4125.5}
	fmt.Println(contaValchan.saldo)
	fmt.Println(contaValchan.sacar(200.00))
	fmt.Println(contaValchan.saldo)
	fmt.Println(contaValchan.sacar(6000.00))

	// contaValchan2 := contaCorrente{"Valchan", 589, 123456, 4125.5}
	// fmt.Println(contaValchan == contaValchan2) //true

	// contaGuilherme := contaCorrente{titular: "Guilherme", numeroAgencia: 980, numeroConta: 345678, saldo: 124.65}
	// fmt.Println(contaGuilherme)

	// var contaCris *contaCorrente
	// contaCris = new(contaCorrente)
	// contaCris.titular = "Cris"
	// contaCris.numeroAgencia = 234
	// fmt.Println(contaCris)

	// var contaCris2 *contaCorrente
	// contaCris2 = new(contaCorrente)
	// contaCris2.titular = "Cris"
	// contaCris2.numeroAgencia = 234
	// fmt.Println(contaCris == contaCris2)   //false pois compara o endereço
	// fmt.Println(*contaCris == *contaCris2) //true pois compara o conteúdo

	// fmt.Println(somando(1))
	// fmt.Println(somando(1,1))
	// fmt.Println(somando(1,1,1))
	// fmt.Println(somando(1,1,2,4))
}
