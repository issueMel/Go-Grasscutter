package kcp

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils/crypto"
	"Go-Grasscutter/utils/lang"
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"sync"
)

func Init() {
	// Kcp Server
	log.SugaredLogger.Info("kcp listens on 22102")
	lis, err := kcp.ListenWithOptions(":22102", nil, 0, 0)
	if err != nil {
		panic(err)
	}
	for {
		// todo limit player number -> kick or limit online number
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		log.SugaredLogger.Info(lang.Translate("messages.game.connect"), conn.RemoteAddr().String())
		sessionManagement.Store(conn, nil)
		go Connected(conn)
	}
}

var sessionManagement sync.Map // *UDPSession, Player

func Connected(conn *kcp.UDPSession) {
	conn.SetNoDelay(1, config.Conf.Server.Game.KcpInterval, 2, 1)
	conn.SetMtu(1400)
	conn.SetWindowSize(256, 256)
	conn.SetACKNoDelay(true)
	// conn.SetDeadline(time.Now().Add(30 * time.Second))
	var buffer = make([]byte, 4096)

	// remove session when error
	defer sessionManagement.Delete(conn)

	for {
		n, e := conn.Read(buffer)
		if e != nil {
			if errors.Is(e, io.ErrClosedPipe) {
				log.SugaredLogger.Info(lang.Translate("messages.game.disconnect"), conn.RemoteAddr())
				return
			}
			log.SugaredLogger.Error(e)
			return
		}
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
			log.SugaredLogger.Error("Read const1 error:", err)
			break
		}
		// Packet sanity check
		if const1 != 17767 {
			// Bad packet
			log.SugaredLogger.Error("Bad Data Package Received: got %d ,expect 17767", const1)
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
			log.SugaredLogger.Error("Read header error:", err)
			break
		}
		_, err = buf.Read(payload)
		if err != nil {
			log.SugaredLogger.Error("Read payload error:", err)
			break
		}

		// Sanity check #2
		err = binary.Read(buf, binary.BigEndian, &const2)
		if err != nil {
			log.SugaredLogger.Error("Read short const2 error:", err)
			break
		}
		if const2 != -30293 {
			log.SugaredLogger.Error("Bad Data Package Received: got %d ,expect -30293", const2)
			break // Bad packet
		}
		// todo Handle
		log.SugaredLogger.Info("handle", opcode)
	}
}
