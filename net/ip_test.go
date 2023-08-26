package net

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	testCase struct {
		ip  net.IP
		ret bool
	}
)

func TestIsIPv4Public(t *testing.T) {
	assertions := assert.New(t)

	cases := []testCase{
		{net.IPv4(192, 168, 0, 1), false},
	}

	for _, c := range cases {
		t.Run(c.ip.String(), func(t *testing.T) {
			r := IsIPv4Public(c.ip)
			assertions.Equal(c.ret, r)
		})
	}

}
