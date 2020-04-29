package client

import (
	"gKV/utils"
	"net"
)

/*
bio send
*/
func Send2Server(operation string, conn *net.TCPConn) {
	//socket write data
	_, err := conn.Write([]byte(operation))
	utils.CheckErr(err)
}

/*
bio receive
*/
func ReceiveFromServer(conn *net.TCPConn) []byte {
	//read data from TCPConn
	var res []byte
	_, err := conn.Read(res)
	utils.CheckErr(err)
	return res
}
