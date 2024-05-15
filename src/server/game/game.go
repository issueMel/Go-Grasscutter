package game

import (
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/game/player"
	"fmt"
	"log"
)

type GameServer struct {
	// Game server base
	Players map[int]*player.Player
	Conn    *kcp.UDPSession

	// Server systems

	// Extra
}

func Init() {
	// Kcp Server
	fmt.Println("kcp listens on 22102")
	lis, err := kcp.ListenWithOptions(":22102", nil, 0, 0)
	if err != nil {
		panic(err)
	}
	conn, e := lis.AcceptKCP()
	if e != nil {
		panic(e)
	}
	conn.SetNoDelay(1, config.Conf.Server.Game.KcpInterval, 2, 1)
	conn.SetMtu(1400)
	// conn.SetDeadline(time.Now().Add(30 * time.Second))
	conn.SetWindowSize(256, 256)
	conn.SetACKNoDelay(true)

	for {
		var buffer = make([]byte, 4096)
		n, e := conn.Read(buffer)
		if e != nil {
			log.Println(e)
			continue
		}
		go handlerReceive(buffer[:n])
	}
}
