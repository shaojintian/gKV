package main

import (
	"fmt"
	"gKV/src/server"
	"gKV/utils"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	//var res string
	//listen port
	netListen, err := net.Listen("tcp", "127.0.0.1:8736")
	utils.CheckErr(err)
	//must close netListen
	defer netListen.Close()

	for {
		fmt.Println("server listening...")
		conn, err := netListen.Accept()
		utils.CheckErr(err)
		server.Handle(conn)
	}
}
