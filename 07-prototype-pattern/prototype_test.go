package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactoryPattern(t *testing.T) {
	cases := []struct {
		name   string
		arg    string
		expect string
	}{
		{"equal1", "Tom", "Tom"},
		{"equal2", "Lucy", "Lucy"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			instance := New(c.arg)
			assert.Equal(t, c.expect, instance.Clone().Description)
		})
	}
}
