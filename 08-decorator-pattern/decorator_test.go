package decorator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoratorPattern(t *testing.T) {
	cases := []struct {
		name     string
		instance Component
		tp       string
		num      int
		expect   Component
	}{
		{name: "equal", instance: &Fruit{Count: 1, Description: "水果"}, tp: "Apple", num: 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.expect = CreateAppleDecorator(c.instance, c.tp, c.num)
			assert.Equal(t, c.expect.GetCount(), c.instance.GetCount()+c.num)
			assert.Equal(t, c.expect.Describe(), fmt.Sprintf("%s, %s", c.instance.Describe(), c.tp))
		})
	}
}
