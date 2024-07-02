package kcp

import (
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	_ "Go-Grasscutter/server/kcp/packet/handler"
	"Go-Grasscutter/server/kcp/session"
	"Go-Grasscutter/utils/lang"
)

func Init() {
	// Kcp Server
	log.SugaredLogger.Info("kcp listens on 22102")
	lis, err := kcp.ListenWithOptions(":22102", nil, 0, 0)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		log.SugaredLogger.Infof(lang.Translate("messages.game.connect"), conn.RemoteAddr().String())
		if s, ok := session.Management.Load(conn); ok {
			go s.(*session.Session).Connected()
		} else {
			sess := session.NewGameSession()
			sess.Tunnel.Kcp = conn
			session.Management.Store(conn, sess)
			go sess.Connected()
		}
	}
}
