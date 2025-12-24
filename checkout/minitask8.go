package checkout

import "fmt"

type Bank struct{}

func (b Bank) setPaymentOne() string {
	return "mandiri"
}

type Online struct{}

func (o Online) setpaymentTwo() string {
	return "qris"
}

type Payment interface {
	PaymentOne() string
	PaymentTwo() string
}

func GetPaymentOne(payment Payment) {
	fmt.Println(payment.PaymentOne())
}

func GetPaymentTwo(payment Payment) {
	fmt.Println(payment.PaymentTwo())
}

func main() {
	
}