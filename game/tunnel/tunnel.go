package tunnel

import (
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/utils/crypto"
)

type Tunnel struct {
	Kcp           *kcp.UDPSession
	LastClientSeq uint32
	LastPingTime  int64
	EncryptKey    []byte
}

func (t *Tunnel) Send(packet *base.Packet) {
	if packet == nil {
		return
	}
	if packet.Opcode <= 0 {
		log.SugaredLogger.Warn("Tried to send packet with missing cmd id!")
		return
	}

	// Header
	if packet.ShouldBuildHeader {
		packet.BuildHeader(t.LastClientSeq)
		t.LastClientSeq++
	}

	bytes := packet.Build()

	if packet.ShouldEncrypt {
		if packet.UseDispatchKey {
			bytes = crypto.Xor(bytes, crypto.DispatchKey)
		} else {
			bytes = crypto.Xor(bytes, t.EncryptKey)
		}
	}

	_, err := t.Kcp.Write(bytes)
	if err != nil {
		log.SugaredLogger.Debug("Unable to send packet to client:", err)
	}
}
