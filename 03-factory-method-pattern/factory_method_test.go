package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePayment(t *testing.T) {
	cases := []struct {
		name    string
		kind    int
		balance float32
		expect  Payment
	}{
		{name: "cash equal", kind: 1, balance: 10.00, expect: &CashPay{10.00}},
		{name: "credit equal", kind: 2, balance: 20.00, expect: &CreditPay{20.00}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			payment, err := GeneratePayment(Kind(c.kind), c.balance)
			assert.Nil(t, err)
			assert.Equal(t, c.expect, payment)
		})
	}
}

func TestPay(t *testing.T) {
	cases := []struct {
		name    string
		kind    int
		balance float32
		money   float32
		expect  float32
	}{
		{name: "cash pay equal", kind: 1, balance: 10.00, money: 2.00, expect: 8.00},
		{name: "credit pay equal", kind: 2, balance: 20.00, money: 2.00, expect: 18.00},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			payment, err := GeneratePayment(Kind(c.kind), c.balance)
			assert.Nil(t, err)
			payment.Pay(c.money)
			assert.Equal(t, c.expect, payment.GetBalance())
		})
	}
}
