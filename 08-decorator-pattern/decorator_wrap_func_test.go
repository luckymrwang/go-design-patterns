package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogDecorate(t *testing.T) {
	cases := []struct {
		name     string
		instance Object
		expect   int
	}{
		{name: "equal", instance: func(n int) int {
			return n * 2
		}, expect: 10},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := LogDecorate(c.instance)
			assert.Equal(t, c.expect, f(5))
		})
	}
}
