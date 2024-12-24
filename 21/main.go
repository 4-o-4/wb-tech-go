package main

import "fmt"

type PaymentAdapter struct {
    oldSystem *OldPaymentSystem
}

func (p *PaymentAdapter) ProcessPayment(amount float64) {
    fmt.Println("Адаптер: преобразование вызова для старой системы")
    p.oldSystem.MakePayment(amount)
}

type PaymentProcessor interface {
    ProcessPayment(amount float64)
}

type NewPaymentSystem struct{}

func (n *NewPaymentSystem) ProcessPayment(amount float64) {
    fmt.Printf("Новая система: обработка платежа на сумму %.2f\n", amount)
}

type OldPaymentSystem struct{}

func (o *OldPaymentSystem) MakePayment(amount float64) {
    fmt.Printf("Старая система: обработка платежа на сумму %.2f\n", amount)
}

func main() {
    newSystem := &NewPaymentSystem{}
    newSystem.ProcessPayment(100.0)

    oldSystem := &OldPaymentSystem{}
    adapter := &PaymentAdapter{oldSystem: oldSystem}
    adapter.ProcessPayment(200.0)
}
