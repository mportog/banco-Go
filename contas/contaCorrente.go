package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso", c.saldo
	}
	return "saldo insuficiente", c.saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Depósito não realizado", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
