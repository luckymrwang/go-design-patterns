package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyPattern(t *testing.T) {
	cases := []struct {
		name   string
		expect int
	}{
		{name: "equal", expect: 20},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NewOrganization().Count()
			assert.Equal(t, c.expect, got)
		})
	}
}
