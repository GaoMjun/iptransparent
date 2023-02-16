package iptransparent

import (
	"bytes"
	"encoding/binary"

	"tcpip"
)

type Header struct {
	Version  tcpip.VERSION
	Protocol tcpip.IPPROTOCOL
	DstIP    []byte
	DstPort  int
}

func (self Header) Marshal() (bs []byte) {
	buffer := &bytes.Buffer{}
	binary.Write(buffer, binary.BigEndian, uint8(self.Version))
	binary.Write(buffer, binary.BigEndian, uint8(self.Protocol))
	binary.Write(buffer, binary.LittleEndian, self.DstIP)
	binary.Write(buffer, binary.BigEndian, uint16(self.DstPort))

	bs = buffer.Bytes()
	return
}
