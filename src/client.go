package src

import (
	"gKV/utils"
	"net"
)

//-------------Global variable-------------------
var Client = newClient()

type gkvClient struct {
	id         int
	operations []byte

	//save operation

}

func newClient() *gkvServer {
	return &gkvServer{
		id:         1024,
		operations: make([]byte, 1024),
	}
}

/*
bio send
*/
func Send2Server(operation string, conn *net.TCPConn) {
	//log.Println("start send operation:"+string(operation)+ " to server... ")
	_, err := conn.Write([]byte(operation))
	utils.CheckErr(err)
	//log.Println("send operation:"+string(operation)+ " to server successfully!")
}

/*
bio receive
*/
func ReceiveFromServer(conn *net.TCPConn) string {
	//log.Println("start receive res FROM server....")
	//read data from TCPConn
	res := make([]byte, 1024)
	n, err := conn.Read(res)
	utils.CheckErr(err)
	//log.Printf(string(res[:n]))
	utils.CheckErr(err)
	//log.Println("receive res:"+string(res[:n])+ " FROM server successfully!")
	return string(res[:n])
}
