package server

import (
	"gKV/utils"
	"net"
)

func ReceiveFromClient(conn net.Conn) []byte {
	operation := make([]byte, 1024)
	_, err := conn.Read(operation)
	utils.CheckErr(err)
	return operation
}

func Send2Client(res string, conn net.Conn) {
	_, err := conn.Write([]byte(res))
	utils.CheckErr(err)
}

func Handle(conn net.Conn) {

	//Receive operation from client
	operation := ReceiveFromClient(conn)
	//Send2Client
	Send2Client(string(operation), conn)
}
