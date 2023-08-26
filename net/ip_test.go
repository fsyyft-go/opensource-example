package net

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	testIPCase struct {
		ip  net.IP
		ret bool
	}
	testNetCase struct {
		ip  net.IP
		net net.IPNet
		ret bool
	}
)

func TestIsIPv4Public(t *testing.T) {
	assertions := assert.New(t)

	cases := []testIPCase{
		{net.IPv4(192, 168, 0, 1), false},
		{net.IPv4(120, 27, 53, 19), true},
	}

	for _, c := range cases {
		t.Run(c.ip.String(), func(t *testing.T) {
			r := IsIPv4Public(c.ip)
			assertions.Equal(c.ret, r)
		})
	}

}

func TestIsIPv4Inner(t *testing.T) {
	assertions := assert.New(t)

	cases := []testIPCase{
		{net.IPv4(192, 168, 0, 1), true},
		{net.IPv4(120, 27, 53, 19), false},
	}

	for _, c := range cases {
		t.Run(c.ip.String(), func(t *testing.T) {
			r := IsIPv4Inner(c.ip)
			assertions.Equal(c.ret, r)
		})
	}
}

func TestIsIPv4(t *testing.T) {
	assertions := assert.New(t)

	cases := []testIPCase{
		{net.IPv4(192, 168, 0, 1), true},
		{net.IPv4(120, 27, 53, 19), true},
	}

	for _, c := range cases {
		t.Run(c.ip.String(), func(t *testing.T) {
			r := IsIPv4(c.ip)
			assertions.Equal(c.ret, r)
		})
	}
}

func TestIsIPv6(t *testing.T) {
	// TODO 补单元测试。
}

func TestContains(t *testing.T) {
	assertions := assert.New(t)

	cases := []testNetCase{
		{net.IPv4(192, 168, 0, 1), inner010, false},
		{net.IPv4(192, 168, 0, 1), inner100064, false},
		{net.IPv4(192, 168, 0, 1), inner172016, false},
		{net.IPv4(192, 168, 0, 1), inner192168, true},
		{net.IPv4(120, 27, 53, 19), inner010, false},
		{net.IPv4(120, 27, 53, 19), inner100064, false},
		{net.IPv4(120, 27, 53, 19), inner172016, false},
		{net.IPv4(120, 27, 53, 19), inner192168, false},
		{net.IPv4(120, 27, 53, 18), inner000, false},
		{net.IPv4(120, 27, 53, 18), inner127, false},
		{net.IPv4(120, 27, 53, 18), inner169254, false},
		{net.IPv4(120, 27, 53, 18), inner192000, false},
		{net.IPv4(120, 27, 53, 18), inner192000002, false},
		{net.IPv4(120, 27, 53, 18), inner192088099, false},
		{net.IPv4(120, 27, 53, 18), inner198018, false},
		{net.IPv4(120, 27, 53, 18), inner198051100, false},
		{net.IPv4(120, 27, 53, 18), inner203000113, false},
		{net.IPv4(120, 27, 53, 18), inner224, false},
		{net.IPv4(120, 27, 53, 18), inner240, false},
		{net.IPv4(120, 27, 53, 18), inner255, false},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%[1]s in %[2]s %[3]v", c.ip.String(), c.net.String(), c.ret)
		t.Run(name, func(t *testing.T) {
			r := c.net.Contains(c.ip)
			assertions.Equal(c.ret, r)
		})
	}
}

func TestIsGlobalUnicast(t *testing.T) {
	assertions := assert.New(t)

	// 全局单播地址参考材料：https://developer.aliyun.com/article/406115。

	cases := []testIPCase{
		{net.IPv4(192, 168, 0, 1), true},
		{net.IPv4(120, 27, 53, 19), true},
	}

	for _, c := range cases {
		t.Run(c.ip.String(), func(t *testing.T) {
			r := c.ip.IsGlobalUnicast()
			assertions.Equal(c.ret, r)
		})
	}
}
