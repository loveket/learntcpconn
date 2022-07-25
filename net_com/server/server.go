package server

import (
	"bufio"
	"log"
	"net"
	"net-com/common"
)

const MaxReadByte = 4096

func ServerStart() {
	log.Println("server start ...")
	lis, err := common.LisConn(common.ServerAddr)
	if err != nil {
		log.Println("lis server addr err=", err)
		return
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		conn.Write([]byte("please send\n"))
		go HandleConn(conn)
	}
}
func HandleConn(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, MaxReadByte)
	bnr := bufio.NewReader(conn)
	for {
		_, err := bnr.Read(b)
		if err != nil {
			log.Println("read from client data err", err)
			return
		}
		log.Println(string(b))
	}

}
