package session

import (
	"Go-Grasscutter/log"
)

var Router = make(map[uint16]func(*Session, []byte, []byte))

func Handle(sess *Session, opcode uint16, header, payload []byte) {
	f, ok := Router[opcode]
	if !ok {
		log.SugaredLogger.Errorf("router %d not exist\n", opcode)
		return
	}
	// todo ENHANCE: fix session logic
	f(sess, header, payload)
}
