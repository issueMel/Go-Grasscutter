package packet

import (
	"Go-Grasscutter/generated/pb"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

type BasePacket struct {
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

func (b BasePacket) BuildHeader(clientSequence int) {
	if b.Header != nil && clientSequence == 0 {
		return
	}
	packetHead := pb.PacketHead{
		ClientSequenceId: uint32(clientSequence),
		SentMs:           uint64(time.Now().UnixMilli()),
	}
	head, err := proto.Marshal(&packetHead)
	if err != nil {
		log.Println(err)
		return
	}
	b.Header = head
}

func (b BasePacket) Build() []byte {
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
