package iptransparent

import (
	"io"
	"net"
	"time"

	"ladder"
	"tcpip"
)

type Server struct {
}

func (self *Server) ServeConn(rwc io.ReadWriteCloser) (err error) {
	var (
		header    Header
		dailer    = &net.Dialer{Timeout: time.Second * 3}
		conn      net.Conn
		remote    *ladder.ConnWithTimeout
		localConn = NewConn(rwc)
	)

	header, err = localConn.ReadHeader()

	if header.Protocol == tcpip.TCP {
		raddr := &net.TCPAddr{IP: header.DstIP, Port: header.DstPort}
		conn, err = dailer.Dial("tcp", raddr.String())
		if err != nil {
			return
		}
	}

	if header.Protocol == tcpip.UDP {
		raddr := &net.UDPAddr{IP: header.DstIP, Port: header.DstPort}
		conn, err = net.DialUDP("udp", nil, raddr)
		if err != nil {
			return
		}
	}

	remote = ladder.NewConnWithTimeout(conn)

	ladder.Pipe(localConn, remote)

	return
}
