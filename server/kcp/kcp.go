package kcp

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils/crypto"
	"bytes"
	"encoding/binary"
)

func Init() {
	// Kcp Server
	log.Info("kcp listens on 22102")
	lis, err := kcp.ListenWithOptions(":22102", nil, 0, 0)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		conn.SetNoDelay(1, config.Conf.Server.Game.KcpInterval, 2, 1)
		conn.SetMtu(1400)
		// conn.SetDeadline(time.Now().Add(30 * time.Second))
		conn.SetWindowSize(256, 256)
		conn.SetACKNoDelay(true)
		var buffer = make([]byte, 4096)
		n, e := conn.Read(buffer)
		if e != nil {
			log.Error(e)
			continue
		}
		// todo limit player number
		go handlerReceive(buffer[:n])
	}
}

func handlerReceive(buffer []byte) {
	// todo use EncryptKey / DispatchKey
	buf := bytes.NewBuffer(crypto.Xor(buffer, crypto.DispatchKey))
	for buf.Len() > 0 {
		if buf.Len() < 12 {
			break
		}
		var const1, const2 int16
		var opcode, headerLength uint16
		var payloadLength uint32
		err := binary.Read(buf, binary.BigEndian, &const1)
		if err != nil {
			log.Error("Read const1 error:", err)
			break
		}
		// Packet sanity check
		if const1 != 17767 {
			// Bad packet
			log.Error("Bad Data Package Received: got %d ,expect 17767", const1)
			break
		}

		// Data
		opcode = binary.BigEndian.Uint16(buf.Next(2))
		headerLength = binary.BigEndian.Uint16(buf.Next(2))
		payloadLength = binary.BigEndian.Uint32(buf.Next(4))
		header := make([]byte, headerLength)
		payload := make([]byte, payloadLength)

		_, err = buf.Read(header)
		if err != nil {
			log.Error("Read header error:", err)
			break
		}
		_, err = buf.Read(payload)
		if err != nil {
			log.Error("Read payload error:", err)
			break
		}

		// Sanity check #2
		err = binary.Read(buf, binary.BigEndian, &const2)
		if err != nil {
			log.Error("Read short const2 error:", err)
			break
		}
		if const2 != -30293 {
			log.Error("Bad Data Package Received: got %d ,expect -30293", const2)
			break // Bad packet
		}
		// todo Handle
		log.Info("handle", opcode)
	}
}
