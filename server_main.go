package main

import (
	"gKV/src/server"
	"gKV/utils"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	var res string
	//listen port
	netListen, err := net.Listen("tcp", "127.0.0.1/50010")
	utils.CheckErr(err)
	//must close netListen
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		utils.CheckErr(err)
		server.Handle(conn)
	}
}
