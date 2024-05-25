package session

import (
	"Go-Grasscutter/log"
)

var Router map[uint16]func(*Session, []byte, []byte) // uint16, func(*Session, []byte, []byte)

func Handle(sess *Session, opcode uint16, header, payload []byte) {
	f, ok := Router[opcode]
	if !ok {
		log.SugaredLogger.Errorf("router %d not exist\n", opcode)
		return
	}
	f(sess, header, payload)
}
