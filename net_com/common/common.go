package common

import (
	"log"
	"net"
)

var (
	ClientAddr = "127.0.0.1:9000"
	ServerAddr = "127.0.0.1:9001"
	MiddleAddr = "127.0.0.1:10000"
)
var MiddleChannel = make(chan net.Conn, 1000)

func LisConn(addr string) (lis *net.TCPListener, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println("addr solv err=", err)
		return nil, err
	}
	lis, err = net.ListenTCP("tcp", tcpAddr)
	return
}
func DailConn(laddr, raddr string) (conn *net.TCPConn, err error) {
	ltcpAddr, err := net.ResolveTCPAddr("tcp", laddr)
	if err != nil {
		log.Println("addr solv err=", err)
		return nil, err
	}
	rtcpAddr, err := net.ResolveTCPAddr("tcp", raddr)
	if err != nil {
		log.Println("addr solv err=", err)
		return nil, err
	}
	conn, err = net.DialTCP("tcp", ltcpAddr, rtcpAddr)
	return
}
