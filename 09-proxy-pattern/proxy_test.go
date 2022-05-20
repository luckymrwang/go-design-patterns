package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyPattern(t *testing.T) {
	cases := []struct {
		name   string
		object *Object
		proxy  *ProxyObject
		arg    string
		expect string
	}{
		{"equal", &Object{Action: "run"}, &ProxyObject{}, "run", "I can run"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, c.proxy.object.ObjDo(c.arg))
		})
	}
}
