package kcp

import (
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/utils/crypto"
	"fmt"
	"github.com/xtaci/kcp-go"
	"io"
	"net"
	"time"
)

// Kcp Server
func KcpInit() {
	fmt.Println("kcp listens on 22102")
	crypt, err := kcp.NewSimpleXORBlockCrypt(crypto.EncryptKey)
	if err != nil {
		return
	}
	lis, err := kcp.ListenWithOptions(":22102", crypt, 0, 0)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		conn.SetNoDelay(1, config.Conf.Server.Game.KcpInterval, 2, 1)
		conn.SetMtu(1400)
		conn.SetDeadline(time.Now().Add(30 * time.Second))
		conn.SetWindowSize(256, 256)
		conn.SetStreamMode(true)
		conn.SetACKNoDelay(true)
		go func(conn net.Conn) {
			var buffer = make([]byte, 1024)
			for {
				n, e := conn.Read(buffer)
				if e != nil {
					if e == io.EOF {
						break
					}
					fmt.Println(e)
					break
				}
				fmt.Printf("data:%v addr:%s count:%v\n", string(buffer[:n]), conn.RemoteAddr(), n)
			}
		}(conn)
	}
}

// UDP Server
func UdpInit() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 22102,
	})
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // receive data
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			continue
		}
		d := crypto.Xor(data[:], crypto.EncryptKey)
		fmt.Printf("data:%v addr:%s count:%v\n", string(d), addr, n)
	}
}
