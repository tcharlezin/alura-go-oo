package features

import (
	"fmt"
	"go-oo/models"
)

func CreateContaCorrente() {

	conta1 := models.ContaCorrente{
		Nome:    "Tcharles",
		Agencia: 1,
		Conta:   "5916650",
		Saldo:   120.50,
	}

	conta2 := models.ContaPoupanca{
		Nome:    "Bruna",
		Agencia: 1,
		Conta:   "5916650",
		Saldo:   45.50,
	}

	conta1.CashIn(100)
	conta1.CashIn(150)
	conta1.CashOut(250)

	status := conta1.Transfer(50, &conta2)

	fmt.Println(status)
	fmt.Println(conta1)
	fmt.Println(conta2)

	status = conta2.Transfer(100, &conta1)

	fmt.Println(status)
	fmt.Println(conta1)
	fmt.Println(conta2)
}
