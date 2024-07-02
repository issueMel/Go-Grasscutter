package session

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils/crypto"
	"Go-Grasscutter/utils/lang"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"sync"
)

var Management sync.Map // *UDPSession, *Session

func (s *Session) Connected() {
	conn := s.Tunnel
	conn.Kcp.SetNoDelay(1, config.Conf.Server.Game.KcpInterval, 2, 1)
	conn.Kcp.SetMtu(1400)
	conn.Kcp.SetWindowSize(256, 256)
	conn.Kcp.SetACKNoDelay(true)
	// conn.SetDeadline(time.Now().Add(30 * time.Second))
	var buffer = make([]byte, 4096)

	// remove session when error
	defer Management.Delete(conn)
	// todo ENHANCE: kick after save()
	//defer func() {
	//	game.Server.KickPlayer(s.Player.ID)
	//	runtime.GC() // todo limit gc times at per second
	//}()
	for {
		n, e := conn.Kcp.Read(buffer)
		if e != nil {
			if errors.Is(e, io.ErrClosedPipe) {
				log.SugaredLogger.Infof(lang.Translate("messages.game.disconnect"), conn.Kcp.RemoteAddr())
				return
			}
			log.SugaredLogger.Error(e)
			return
		}
		go s.handleReceive(buffer[:n])
	}
}

func (s *Session) handleReceive(buffer []byte) {
	var buf *bytes.Buffer
	// Decrypt and turn back into a packet
	if s.UseSecretKey {
		buf = bytes.NewBuffer(crypto.Xor(buffer, s.Tunnel.EncryptKey))
	} else {
		buf = bytes.NewBuffer(crypto.Xor(buffer, crypto.DispatchKey))
	}
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
			log.SugaredLogger.Errorf("Bad Data Package Received: got %d ,expect 17767", const1)
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
			log.SugaredLogger.Errorf("Bad Data Package Received: got %d ,expect -30293", const2)
			break // Bad packet
		}
		// Handle
		if _, ok := dontHandle[int(opcode)]; ok {
			return
		}
		log.SugaredLogger.Info("handle: ", opcode)
		Handle(s, opcode, header, payload)
	}
}

var dontHandle = make(map[int]struct{})

func init() {
	add(27557, 20100, 6144, 2030)
}
func add(num ...int) {
	for _, n := range num {
		dontHandle[n] = struct{}{}
	}
}
