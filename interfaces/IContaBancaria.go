package interfaces

type IContaBancaria interface {
	CashOut(valor float32) string
	CashIn(valor float32) (string, float32)
	Transfer(valorTransferencia float32, contaDestino IContaBancaria) string
}
