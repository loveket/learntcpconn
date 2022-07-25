package transconn

import (
	"io"
	"log"
	"net"
	"net-com/common"
)

func MiddleConn() {
	log.Println("middle start ...")
	mLis, err := common.LisConn(common.MiddleAddr)
	if err != nil {
		log.Println("middle lis err", err)
		return
	}
	for {
		clientconn, err := mLis.AcceptTCP()
		if err != nil {
			log.Println("mLis err", err)
			return
		}
		saddr, err := net.ResolveTCPAddr("tcp", common.ServerAddr)
		if err != nil {
			return
		}
		serverconn, err := net.DialTCP("tcp", nil, saddr)
		if err != nil {
			log.Println("conn server err", err)
			return
		}
		go io.Copy(serverconn, clientconn)
		go io.Copy(clientconn, serverconn)
	}

}
