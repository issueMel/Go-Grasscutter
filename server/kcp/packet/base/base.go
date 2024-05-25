package base

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"bytes"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"time"
)

type Packet struct {
	const1            int
	const2            int
	ShouldEncrypt     bool
	Opcode            int
	ShouldBuildHeader bool
	Header            []byte
	Data              []byte
	// Encryption
	UseDispatchKey bool
}

func (b *Packet) BuildHeader(clientSequence uint32) {
	if len(b.Header) > 0 && clientSequence == 0 {
		return
	}
	packetHead := pb.PacketHead{
		ClientSequenceId: clientSequence,
		SentMs:           uint64(time.Now().UnixMilli()),
	}
	head, err := proto.Marshal(&packetHead)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	b.Header = head
}

func (b *Packet) Build() []byte {
	b.const1 = 17767
	b.const2 = -30293
	b.ShouldEncrypt = true

	l := 2 + 2 + 2 + 4 + len(b.Header) + len(b.Data) + 2
	buf := bytes.NewBuffer(make([]byte, 0, l))

	binary.Write(buf, binary.BigEndian, uint16(b.const1))
	binary.Write(buf, binary.BigEndian, uint16(b.Opcode))
	binary.Write(buf, binary.BigEndian, uint16(len(b.Header)))
	binary.Write(buf, binary.BigEndian, uint32(len(b.Data)))
	buf.Write(b.Header)
	buf.Write(b.Data)
	binary.Write(buf, binary.BigEndian, uint16(b.const2))

	return buf.Bytes()
}
