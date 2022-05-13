package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactoryPattern(t *testing.T) {
	cases := []struct {
		name    string
		factory Factory
		arg     string
		expect  Product
	}{
		{"equal1", &ConCreteFactory{}, "Tom", &ConcreteProduct{Name: "Tom"}},
		{"equal2", &ConCreteFactory{}, "Lucy", &ConcreteProduct{Name: "Lucy"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			product := c.factory.CreateProduct(c.arg)
			assert.Equal(t, c.expect, product)
		})
	}
}
