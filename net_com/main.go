package main

import (
	"net-com/client"
	"net-com/server"
	"net-com/transconn"
	"strconv"
	"time"
)

//var wg sync.WaitGroup
func main() {
	//wg.Add(2)
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	return
	// }
	// trace.Start(f)
	// defer f.Close()
	go server.ServerStart()
	time.Sleep(3 * time.Second)
	go transconn.MiddleConn()
	time.Sleep(3 * time.Second)
	for i := 1; i <= 1000; i++ {
		//go client.ClientStart(strconv.Itoa(i))
		go client.ClientToMiddleStart(strconv.Itoa(i))
	}
	//time.Sleep(180 * time.Second)
	//pprof.WriteHeapProfile(f)
	//defer trace.Stop()
	select {}
}
