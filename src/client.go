package src

import (
	"gKV/utils"
	"net"
)

/*
bio send
*/
func Send2Server(operation string, conn *net.TCPConn) {
	//socket write data
	//log.Println("operation: "+operation)
	_, err := conn.Write([]byte(operation))
	utils.CheckErr(err)
}

/*
bio receive
*/
func ReceiveFromServer(conn *net.TCPConn) string {
	//read data from TCPConn
	res := make([]byte, 1024)
	n, err := conn.Read(res)
	utils.CheckErr(err)
	//log.Printf(string(res[:n]))
	utils.CheckErr(err)
	return string(res[:n])
}
