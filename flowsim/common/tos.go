package common

import (
	// "fmt"
	"net"
	"os"
	"syscall"
)

func SetTcpTos(Conn *net.TCPConn, tos int) error {
	f, err := Conn.File()
	if err != nil {
		return err
	}
	host, _, _ := net.SplitHostPort(Conn.LocalAddr().String())
	// fmt.Printf("Local host is: %s\n", host)

	ip, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return err
	}
	// fmt.Printf("Check IPv6 %v\n", ip.IP.To4() == nil)
	return SetTos(f, tos, ip.IP.To4() == nil)
}

func SetUdpTos(conn *net.UDPConn, tos int, isIpv6 bool) error {

	f, err := conn.File()
	if err != nil {
		return err
	}
	// fmt.Printf("Check IPv6 %v\n", isIpv6)
	return SetTos(f, tos, isIpv6)
}

func SetTos(f *os.File, tos int, ipv6 bool) error {

	proto := syscall.IPPROTO_IP
	call := syscall.IP_TOS
	errmsg := "TOS"

	if ipv6 {
		proto = syscall.IPPROTO_IPV6
		call = syscall.IPV6_TCLASS
		errmsg = "TCLASS"
	}
	// fmt.Printf("Setting %s to 0x%x\n", errmsg, tos)

	err := syscall.SetsockoptInt(int(f.Fd()), proto, call, tos)

	if err != nil {
		WarnErrorf(err, "while setting %s to %02x", errmsg, tos)
	}

	return err
}
