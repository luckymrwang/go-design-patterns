package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {
	cases := []struct {
		name    string
		builder *Car
		expect  Builder
	}{
		{"equal", &Car{Vehicle{Wheels: 4, Seats: 4, Structure: "Car"}}, &Car{Vehicle{Wheels: 4, Seats: 4, Structure: "Car"}}},
		{"wheel count", &Car{Vehicle{Wheels: 3, Seats: 4, Structure: "Car"}}, &Car{Vehicle{Wheels: 4, Seats: 4, Structure: "Car"}}},
		{"structure name", &Car{Vehicle{Wheels: 4, Seats: 4, Structure: "bus"}}, &Car{Vehicle{Wheels: 4, Seats: 4, Structure: "Car"}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.builder.Construct()
			assert.Equal(t, c.expect, c.builder)
		})
	}
}
