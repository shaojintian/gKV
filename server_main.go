package main

import (
	"fmt"
	"gKV/src"
	"gKV/utils"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	//init global map
	src.GlobalMap = make(map[string]string, src.MAP_INIT_SIZE)
}

func main() {

	//----------monitor signal to exit gracefully----------
	sigC := make(chan os.Signal)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGSTOP)
	go src.GracefullyExit(sigC)

	//----------start server-------------------------------
	//listen port
	netListen, err := net.Listen("tcp", "127.0.0.1:8737")
	utils.CheckErr(err)
	//only one conn
	conn, err := netListen.Accept()
	utils.CheckErr(err)
	//must close netListen
	defer netListen.Close()

	for {
		fmt.Println("server listening...")
		src.Handle(conn)
	}
}
