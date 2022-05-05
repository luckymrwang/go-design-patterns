package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonPattern(t *testing.T) {
	cases := []struct {
		name     string
		instance singleton
		arg1     string
		arg2     string
		expect   singleton
	}{
		{name: "equal", arg1: "Tom", arg2: "Jack"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.instance = New()
			c.instance["name"] = c.arg1
			c.expect = New()
			assert.Equal(t, c.expect, c.instance)

			// change name
			c.expect["name"] = c.arg2
			assert.Equal(t, c.expect, c.instance)
		})
	}
}
