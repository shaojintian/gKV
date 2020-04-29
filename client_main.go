package main

import (
	"bufio"
	"fmt"
	"gKV/src/client"
	"gKV/utils"
	"log"
	"net"
	"os"
	"strings"
)

//socket client based on TCP

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	//remote tcpAddr(server addr)
	addr := "127.0.0.1/50010"
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
	//127.0.0.1:xxxx> set k v
	for {
		fmt.Print(addr + "> ")
		operation, err := oreader.ReadString('\n')
		utils.CheckErr(err)
		//delete '\n' in operation
		operation = strings.Replace(operation, "\n", "", -1)
		//send operation
		client.Send2Server(operation, conn)
		res := client.ReceiveFromServer(conn)
		if len(res) == 0 {
			fmt.Println(utils.ERR_NIL)
		} else {
			fmt.Println(res)
		}
	}

}
