package main

import (
	"bufio"
	"fmt"
	"gKV/src"
	"gKV/utils"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//socket client based on TCP
//
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	//----------monitor signal to exit gracefully----------
	sigC := make(chan os.Signal)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGSTOP)
	go src.GracefullyExit(sigC)

	//----------start client---------------
	//remote tcpAddr(server addr)
	addr := "127.0.0.1:8736"
	//read  [operator  key  value] from commandline
	oreader := bufio.NewReader(os.Stdin)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	utils.CheckErr(err)
	//establish tcp connection
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	utils.CheckErr(err)
	//must close conn
	defer conn.Close()

	//send data to Server
	//like
	for {
		//127.0.0.1:xxxx> set k v
		fmt.Print(addr + "> ")
		operation, err := oreader.ReadString('\n')
		utils.CheckErr(err)
		//delete '\n' in operation
		operation = strings.Replace(operation, "\n", "", -1)
		//send operation
		src.Send2Server(operation, conn)
		res := src.ReceiveFromServer(conn)
		if len(res) == 0 {
			fmt.Println(utils.ERR_NIL)
		} else {
			fmt.Println(res)
		}
	}

}
