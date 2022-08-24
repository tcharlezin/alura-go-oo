package models

import (
	"go-oo/interfaces"
	"gorm.io/gorm"
)

type ContaPoupanca struct {
	gorm.Model
	Nome    string
	Agencia int
	Conta   string
	Saldo   float32
}

func (c *ContaPoupanca) CashOut(valor float32) string {

	if valor < 0 {
		return "Valor inválido!"
	}

	podeSacar := c.Saldo > valor

	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso!"
	}

	return "Saque indisponível!"
}

func (c *ContaPoupanca) CashIn(valor float32) (string, float32) {

	if valor < 0 {
		return "Valor inválido!", c.Saldo
	}

	podeSacar := valor > 0

	if podeSacar {
		c.Saldo += valor
		return "Deposito realizado com sucesso!", c.Saldo
	}

	return "Deposito indisponível!", c.Saldo
}

func (c *ContaPoupanca) Transfer(valorTransferencia float32, contaDestino interfaces.IContaBancaria) string {
	if c.Saldo < valorTransferencia || valorTransferencia <= 0 {
		return "Saldo insuficiente!"
	}

	c.CashOut(valorTransferencia)
	contaDestino.CashIn(valorTransferencia)

	return "Transferência realizada com sucesso!"
}
