package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratorPattern(t *testing.T) {
	cases := []struct {
		name   string
		start  int
		end    int
		expect int
	}{
		{name: "equal", start: 1, end: 20, expect: 20},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			count := 0
			for range Count(c.start, c.end) {
				count++
			}
			assert.Equal(t, c.expect, count)
		})
	}
}
