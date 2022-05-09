package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectPoolPattern(t *testing.T) {
	cases := []struct {
		name   string
		count  int
		expect int
	}{
		{name: "equal", count: 2, expect: 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pool := NewPool(c.count)
			assert.Equal(t, c.expect, len(*pool))

			for p := range *pool {
				p.Do()
			}
		})
	}
}
