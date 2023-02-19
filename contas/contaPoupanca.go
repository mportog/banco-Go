package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso", c.saldo
	}
	return "saldo insuficiente", c.saldo
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Depósito não realizado", c.saldo
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
