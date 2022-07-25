package client

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const MaxReadByte = 4096

func ClientStart(index string) {
	log.Println(index, "client start ...")
	// conn, err := common.DailConn(common.ClientAddr, common.ServerAddr)
	// if err != nil {
	// 	log.Println("conn server err", err)
	// 	return
	// }
	conn, err := net.Dial("tcp", ":9001")
	if err != nil {
		log.Println("conn server err", err)
		return
	}
	b := make([]byte, MaxReadByte)
	conn.Read(b[:])
	res := strings.Split(string(b), "\n")
	switch res[0] {
	case "please send":
		ticker := time.NewTicker(3 * time.Second)
		i := 0
		go func() {
			for v := range ticker.C {
				i++
				conn.Write([]byte("*" + index + "*" + strconv.Itoa(i) + "---" + v.Format("2006-01-02 15:04:05")))
			}
		}()

	}
}
func ClientToMiddleStart(index string) {
	log.Println(index, "client start ...")
	// conn, err := common.DailConn(common.ClientAddr, common.ServerAddr)
	// if err != nil {
	// 	log.Println("conn server err", err)
	// 	return
	// }
	conn, err := net.Dial("tcp", ":10000")
	if err != nil {
		log.Println("conn server err", err)
		return
	}
	b := make([]byte, MaxReadByte)
	conn.Read(b[:])
	res := strings.Split(string(b), "\n")
	switch res[0] {
	case "please send":
		ticker := time.NewTicker(3 * time.Second)
		i := 0
		go func() {
			for v := range ticker.C {
				i++
				conn.Write([]byte("*" + index + "*" + strconv.Itoa(i) + "---" + v.Format("2006-01-02 15:04:05")))
			}
		}()

	}
}
