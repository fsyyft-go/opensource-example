package net

import (
	"net"
)

var (
	inner000       = net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.CIDRMask(8, 32)}          // RFC 5735 0.0.0.0/8 本网络（仅作为源地址时合法）。
	inner010       = net.IPNet{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)}         // RFC 1918 10.0.0.0/8 专用网络。
	inner100064    = net.IPNet{IP: net.IPv4(100, 64, 0, 0), Mask: net.CIDRMask(10, 32)}      // RFC 6598 100.64.0.0/10 电信级NAT。
	inner127       = net.IPNet{IP: net.IPv4(127, 0, 0, 0), Mask: net.CIDRMask(8, 32)}        // RFC 5735 127.0.0.0/8 环回。
	inner169254    = net.IPNet{IP: net.IPv4(169, 254, 0, 0), Mask: net.CIDRMask(16, 32)}     // RFC 3927 169.254.0.0/16 链路本地。
	inner172016    = net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)}      // RFC 1918 172.16.0.0/12 专用网络。
	inner192000    = net.IPNet{IP: net.IPv4(192, 0, 0, 0), Mask: net.CIDRMask(24, 32)}       // RFC 5735 192.0.0.0/24 保留（IANA）。
	inner192000002 = net.IPNet{IP: net.IPv4(192, 0, 2, 0), Mask: net.CIDRMask(24, 32)}       // RFC 5735 192.0.2.0/24 TEST-NET-1，文档和示例。
	inner192088099 = net.IPNet{IP: net.IPv4(192, 88, 99, 0), Mask: net.CIDRMask(24, 32)}     // RFC 3068 192.88.99.0/24 6to4中继。
	inner192168    = net.IPNet{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)}     // RFC 1918 192.168.0.0/16 专用网络。
	inner198018    = net.IPNet{IP: net.IPv4(198, 18, 0, 0), Mask: net.CIDRMask(15, 32)}      // RFC 2544 198.18.0.0/15 网络基准测试。
	inner198051100 = net.IPNet{IP: net.IPv4(198, 51, 100, 0), Mask: net.CIDRMask(24, 32)}    // RFC 5737 198.51.100.0/24 TEST-NET-2，文档和示例。
	inner203000113 = net.IPNet{IP: net.IPv4(203, 0, 113, 0), Mask: net.CIDRMask(24, 32)}     // RFC 5737 203.0.113.0/24 TEST-NET-3，文档和示例。
	inner224       = net.IPNet{IP: net.IPv4(224, 0, 113, 0), Mask: net.CIDRMask(4, 32)}      // RFC 3171 224.0.0.0/4 多播（之前的D类网络）。
	inner240       = net.IPNet{IP: net.IPv4(240, 0, 113, 0), Mask: net.CIDRMask(4, 32)}      // RFC 1700 240.0.0.0/4 保留（之前的E类网络）。
	inner255       = net.IPNet{IP: net.IPv4(255, 255, 255, 255), Mask: net.CIDRMask(32, 32)} // RFC 919 255.255.255.255 受限广播。
)

// IPExtend 对 net.IP 进行扩展。
type IPExtend struct {
	*net.IP
}

// IsIPv4Public 是否公网地址。
func (e *IPExtend) IsIPv4Public() bool {
	isIPv4Public := e.IsIPv4() && !e.IsIPv4Inner() &&
		!(inner000.Contains(*e.IP) || inner127.Contains(*e.IP) || inner169254.Contains(*e.IP)) &&
		!(inner192000.Contains(*e.IP) || inner192000002.Contains(*e.IP) || inner192088099.Contains(*e.IP)) &&
		!(inner198018.Contains(*e.IP) || inner198051100.Contains(*e.IP) || inner203000113.Contains(*e.IP)) &&
		!(inner224.Contains(*e.IP) || inner240.Contains(*e.IP) || inner255.Contains(*e.IP))
	return isIPv4Public
}

// IsIPv4Inner 是否是内网地址，不包含回环地址其它测试地址。
func (e *IPExtend) IsIPv4Inner() bool {
	isIPv4Inner := e.IsIPv4() && (inner010.Contains(*e.IP) || inner100064.Contains(*e.IP) || inner172016.Contains(*e.IP) || inner192168.Contains(*e.IP))
	return isIPv4Inner
}

// IsIPv4 是否是 IPv4 格式的地址。
func (e *IPExtend) IsIPv4() bool {
	return nil != e.IP.To4()
}

// IsIPv6 是否是 IPv6 格式的地址。
func (e *IPExtend) IsIPv6() bool {
	return nil != e.IP.To16()
}

// IsIPv4Public 是否公网地址。
func IsIPv4Public(ip net.IP) bool {
	return (&IPExtend{IP: &ip}).IsIPv4Public()
}

// IsIPv4Inner 是否是内网地址，不包含回环地址其它测试地址。
func IsIPv4Inner(ip net.IP) bool {
	return (&IPExtend{IP: &ip}).IsIPv4Inner()
}

// IsIPv4 是否是 IPv4 格式的地址。
func IsIPv4(ip net.IP) bool {
	return (&IPExtend{IP: &ip}).IsIPv4()
}

// IsIPv6 是否是 IPv6 格式的地址。
func IsIPv6(ip net.IP) bool {
	return (&IPExtend{IP: &ip}).IsIPv6()
}
