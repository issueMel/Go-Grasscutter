package session

import (
	"Go-Grasscutter/log"
	"sync"
)

var Router sync.Map // uint16, func(*Session, []byte, []byte)

func Handle(sess *Session, opcode uint16, header, payload []byte) {
	val, ok := Router.Load(opcode)
	if !ok {
		log.SugaredLogger.Errorf("router %d not exist\n", opcode)
		return
	}
	f := val.(func(*Session, []byte, []byte))
	f(sess, header, payload)
}
