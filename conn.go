package iptransparent

import (
	"errors"
	"fmt"
	"io"

	"tcpip"
)

type Conn struct {
	rwc io.ReadWriteCloser
}

func NewConn(rwc io.ReadWriteCloser) (conn *Conn) {
	conn = &Conn{}
	conn.rwc = rwc
	return
}

func (self *Conn) Write(p []byte) (n int, err error) {
	n, err = self.rwc.Write(p)
	return
}

func (self *Conn) Read(p []byte) (n int, err error) {
	n, err = self.rwc.Read(p)
	return
}

func (self *Conn) Close() (err error) {
	err = self.rwc.Close()
	return
}

func (self *Conn) WriteHeader(p Header) (err error) {
	_, err = self.rwc.Write(p.Marshal())
	return
}

func (self *Conn) ReadHeader() (p Header, err error) {
	h := make([]byte, 1)
	_, err = io.ReadFull(self.rwc, h)
	if err != nil {
		return
	}
	p.Version = tcpip.VERSION(h[0])
	if p.Version != tcpip.IPv4 && p.Version != tcpip.IPv6 {
		err = errors.New(fmt.Sprint("version is invalid", p.Version))
		return
	}

	_, err = io.ReadFull(self.rwc, h)
	if err != nil {
		return
	}
	p.Protocol = tcpip.IPPROTOCOL(h[0])
	if p.Protocol != tcpip.TCP && p.Protocol != tcpip.UDP && p.Protocol != tcpip.ICMP {
		err = errors.New(fmt.Sprint("protocol is invalid", p.Protocol))
		return
	}

	if p.Version == tcpip.IPv4 {
		h = make([]byte, 4)
		_, err = io.ReadFull(self.rwc, h)
		if err != nil {
			return
		}
		p.DstIP = h
	}

	if p.Version == tcpip.IPv6 {
		h = make([]byte, 16)
		_, err = io.ReadFull(self.rwc, h)
		if err != nil {
			return
		}
		p.DstIP = h
	}

	h = make([]byte, 2)
	_, err = io.ReadFull(self.rwc, h)
	if err != nil {
		return
	}
	p.DstPort = int(h[0])<<8 | int(h[1])

	return
}
