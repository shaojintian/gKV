package src

import (
	"gKV/utils"
	"log"
	"net"
	"strings"
)

func ReceiveFromClient(conn net.Conn) []byte {
	log.Println("start read data from client ...")
	operation := make([]byte, 1024)
	_, err := conn.Read(operation)
	utils.CheckErr(err)
	log.Println("read operation:" + string(operation) + " FROM client successfully!")
	return operation
}

func Send2Client(res string, conn net.Conn) {
	log.Println("start send " + res + " to client...!")
	_, err := conn.Write([]byte(res))
	utils.CheckErr(err)
	log.Println("return " + res + " to client successfully!\n")
}

func Handle(conn net.Conn) {

	//Receive operation from client
	operation := ReceiveFromClient(conn)
	//handle xxx k v
	res := doHandle(operation)
	//Send2Client
	Send2Client(res, conn)
}

func doHandle(operation []byte) string {
	var res string
	//handle set/get
	//delete space in operation begin or end
	opStr := strings.TrimSpace(string(operation))
	//split opStr set k v
	opts := strings.Split(opStr, " ")
	//
	log.Println(opts)
	if len(opts) == 3 {
		if opts[0] == "set" {
			res = doSet(opts[1:])
		}
	} else if len(opts) == 2 {
		if opts[0] == "get" {
			res = doGet(opts)
		}
	} else if opts[0] == "size" {
		res = string(zMapSize())

	} else {
		res = "(error) ERR wrong number of arguments for  command"
	}

	return res
}

func doSet(opts []string) string {
	GlobalMap[opts[0]] = opts[1]
	return utils.OK
}

func doGet(opts []string) string {
	v, ok := GlobalMap[opts[1]]
	if ok {
		return v
	}
	return utils.ERR_NIL
}