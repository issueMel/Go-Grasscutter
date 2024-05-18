package kcp

import (
	"Go-Grasscutter/log"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"net"
	"sync"
)

type handshakeWaiter struct {
	ConvId uint64
	Addr   string
}

var (
	handshakeWaiters sync.Map
)

func handleEnet(data []byte, l *Listener, addr net.Addr) {
	if len(data) != 20 {
		return
	}
	var code, enet int32

	buffer := bytes.NewBuffer(data)
	binary.Read(buffer, binary.BigEndian, &code)
	buffer.Next(4) // Empty
	buffer.Next(4) // Empty
	binary.Read(buffer, binary.BigEndian, &enet)
	buffer.Next(4) // Empty

	convID, err := rand.Int(rand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		// Handle error
		panic(err)
	}
	switch code {
	case 255:
		sendHandshakeRsp(enet, l, addr, convID.Int64())
	case 404:
		if l != nil {
			l.closeSession(addr)
		}
	}
}

func sendHandshakeRsp(enet int32, l *Listener, addr net.Addr, conv int64) {
	packet := new(bytes.Buffer)
	binary.Write(packet, binary.BigEndian, int32(325))
	binary.Write(packet, binary.LittleEndian, int32(conv>>32))
	binary.Write(packet, binary.LittleEndian, int32(conv&int64(4294967295)))
	binary.Write(packet, binary.BigEndian, int32(enet))
	binary.Write(packet, binary.BigEndian, int32(340870469))
	udpSend(packet.Bytes(), l, addr)
}

func sendDisconnectPacket(conv int64, code int, l *Listener, addr net.Addr) {
	packet := new(bytes.Buffer)
	binary.Write(packet, binary.BigEndian, int32(404))
	binary.Write(packet, binary.LittleEndian, int32(conv>>32))
	binary.Write(packet, binary.LittleEndian, int32(conv&int64(4294967295)))
	binary.Write(packet, binary.BigEndian, int32(code))
	binary.Write(packet, binary.BigEndian, int32(423728276))
	udpSend(packet.Bytes(), l, addr)
}

func udpSend(data []byte, s *Listener, addr net.Addr) {
	_, err := s.conn.WriteTo(data, addr)
	if err != nil {
		log.Error(err)
		return
	}
}
