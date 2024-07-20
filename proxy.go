package proxy

import "net"

const Network = "tcp"

var localAddress *net.TCPAddr = &net.TCPAddr{IP: net.IP{127, 0, 0, 1},
	Port: 1080,
	Zone: "",
}

func LocalAddress() *net.TCPAddr {
	return localAddress
}
