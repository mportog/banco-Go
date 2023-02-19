package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	contaPorto := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Matheus Porto",
			Cpf:       "123.123.123-12",
			Profissao: "Desenvolvedor"},
		Agencia: 2132,
		Conta:   4,
	}
	contaPorto.Depositar(32.56)

	contaGuilherme := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Guilherme Alura",
			Cpf:       "321.321.321-32",
			Profissao: "Professor"},
		Agencia: 2332,
		Conta:   3}
	contaGuilherme.Depositar(232.24)

	fmt.Println(contaPorto)
	fmt.Println(contaGuilherme)

	clienteCris := clientes.Titular{"Cristiane", "123.456.789-10", "A maior de todas"}
	var contaCris *contas.ContaCorrente
	contaCris = new(contas.ContaCorrente)
	contaCris.Titular = clienteCris
	contaCris.Depositar(500.32)

	fmt.Println(*contaCris)
	fmt.Println(contaCris.Saldo())

	contaPorto.Sacar(15.05)
	fmt.Println(contaPorto)

	status, valor := contaCris.Depositar(200)
	fmt.Println(status, valor)

	status, valor = contaCris.Sacar(55.50)
	fmt.Println(status, valor)

	contaGuilherme.Transferir(120.35, &contaPorto)

	fmt.Println(contaGuilherme.Saldo())
	pagarBoleto(&contaGuilherme, 40)
	fmt.Println(contaGuilherme.Saldo())

}
