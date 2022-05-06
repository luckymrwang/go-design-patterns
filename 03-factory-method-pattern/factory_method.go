package factory

import (
	"errors"
)

/*
	factory method design pattern will create object with the exact type
	设计思想：
		类型常量
		接口factory
		生成函数
		实现接口方法的struct
*/

type Kind int

const (
	Cash Kind = 1 << iota
	Credit
)

type Payment interface {
	Pay(money float32) error
	GetBalance() float32
}

// 实现两个struct，继承接口Payment
type CashPay struct {
	Balance float32
}

func (cash *CashPay) Pay(money float32) error {
	if cash.Balance < 0 || cash.Balance < money {
		return errors.New("balance not enough")
	}
	cash.Balance -= money
	return nil
}

func (cash *CashPay) GetBalance() float32 {
	return cash.Balance
}

type CreditPay struct {
	Balance float32
}

func (credit *CreditPay) Pay(money float32) error {
	if credit.Balance < 0 || credit.Balance < money {
		return errors.New("balance not enough")
	}
	credit.Balance -= money
	return nil
}

func (credit *CreditPay) GetBalance() float32 {
	return credit.Balance
}

// factory method pattern
func GeneratePayment(k Kind, balance float32) (Payment, error) {
	switch k {
	case Cash:
		return &CashPay{Balance: balance}, nil
	case Credit:
		return &CreditPay{Balance: balance}, nil
	default:
	}

	return nil, errors.New("Payment do not support this")
}
